Feature: Test internal and external twap calculation

    Background:
        # epoch time is 1602806400
        Given time is updated to "2020-10-16T00:00:00Z"
    And the following assets are registered:
            | id         | decimal places | quantum |
            | USD.3.1000 | 3              | 1000    |
    And the perpetual oracles from "0xCAFECAFE1":
            | name        | asset      | settlement property | settlement type | schedule property | schedule type  | margin funding factor | interest rate | clamp lower bound | clamp upper bound | quote name | settlement decimals |
            | perp-oracle | USD.3.1000 | perp.ETH.value      | TYPE_INTEGER    | perp.funding.cue  | TYPE_TIMESTAMP | 0.5                   | 0.00          | 0.0               | 0.0               | ETH        | 18                  |
    And the liquidity sla params named "SLA":
            | price range | commitment min time fraction | performance hysteresis epochs | sla competition factor |
            | 100.0       | 0.5                          | 1                             | 1.0                    |
    And the price monitoring named "my-price-monitoring":
            | horizon | probability | auction extension |
            | 43200   | 0.9999999   | 120               |
    And the log normal risk model named "my-log-normal-risk-model":
            | risk aversion | tau                    | mu | r     | sigma |
            | 0.000001      | 0.00011407711613050422 | 0  | 0.016 | 0.8   |
    And the markets:
            | id        | quote name | asset      | risk model               | margin calculator         | auction duration | fees         | price monitoring    | data source config | linear slippage factor | quadratic slippage factor | position decimal places | market type | sla params |
            | ETH/DEC19 | ETH        | USD.3.1000 | my-log-normal-risk-model | default-margin-calculator | 120              | default-none | my-price-monitoring | perp-oracle        | 0.25                   | 0                         | -3                      | perp        | SLA        |
    And the following network parameters are set:
            | name                                    | value |
            | network.markPriceUpdateMaximumFrequency | 0s    |
    And the average block duration is "1"
        When the parties deposit on asset's general account the following amount:
            | party  | asset      | amount          |
            | party1 | USD.3.1000 | 100000000000000 |
            | party2 | USD.3.1000 | 100000000000000 |
            | party3 | USD.3.1000 | 100000000000000 |
            | aux    | USD.3.1000 | 100000000000000 |
            | aux2   | USD.3.1000 | 100000000000000 |
            | lpprov | USD.3.1000 | 100000000000000 |
        Then the parties submit the following liquidity provision:
            | id  | party  | market id | commitment amount | fee   | lp type    |
            | lp1 | lpprov | ETH/DEC19 | 100000000         | 0.001 | submission |
    # move market to continuous
    And the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     |
            | aux2   | ETH/DEC19 | buy  | 1      | 1     | 0                | TYPE_LIMIT | TIF_GTC |
            | lpprov | ETH/DEC19 | buy  | 100000 | 1     | 0                | TYPE_LIMIT | TIF_GTC |
            | aux2   | ETH/DEC19 | buy  | 1      | 10    | 0                | TYPE_LIMIT | TIF_GTC |
            | aux    | ETH/DEC19 | sell | 1      | 10    | 0                | TYPE_LIMIT | TIF_GTC |
            | aux    | ETH/DEC19 | sell | 1      | 35    | 0                | TYPE_LIMIT | TIF_GTC |
            | lpprov | ETH/DEC19 | sell | 100000 | 35    | 0                | TYPE_LIMIT | TIF_GTC |
    And the market data for the market "ETH/DEC19" should be:
            | target stake | supplied stake |
            | 4315000      | 100000000      |
    And the opening auction period ends for market "ETH/DEC19"

    @Perpetual @twap
    Scenario: 0053-PERP-028 Internal and External TWAP calculation, auction in funding period
        Given the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/DEC19"
    And time is updated to "2020-10-16T00:05:00Z"
        When the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     |
            | party1 | ETH/DEC19 | buy  | 1      | 10    | 0                | TYPE_LIMIT | TIF_GTC |
            | party2 | ETH/DEC19 | sell | 1      | 10    | 1                | TYPE_LIMIT | TIF_GTC |
        Then time is updated to "2020-10-16T00:10:00Z"

        # 0 min in to the next funding period
        Given the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name             | value                | time offset |
            | perp.ETH.value   | 11000000000000000000 | -1s         |
            | perp.funding.cue | 1602807000           | 0s          |
    And the mark price should be "10" for the market "ETH/DEC19"

        # 1 min in to the next funding period
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10000         | 11000         | -1000           |
        Given the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     |
            | party1 | ETH/DEC19 | buy  | 1      | 11    | 0                | TYPE_LIMIT | TIF_GTC |
            | party2 | ETH/DEC19 | sell | 1      | 11    | 1                | TYPE_LIMIT | TIF_GTC |
    And the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value               | time offset |
            | perp.ETH.value | 9000000000000000000 | 0s          |

        # 2 min in to the next funding period
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10500         | 10000         | 500             |

        # 3 min in to the next funding period
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10666         | 9666          | 1000            |
        Given the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value                | time offset |
            | perp.ETH.value | 10000000000000000000 | 0s          |

        # 4 min in to the next funding period
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10750         | 9750          | 1000            |

        # 5 min in to the next funding period, the auction period will start
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10800         | 9800          | 1000            |
        Given the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     | reference        |
            | aux2   | ETH/DEC19 | buy  | 1      | 15    | 0                | TYPE_LIMIT | TIF_GTC | trigger-auction2 |
            | aux    | ETH/DEC19 | sell | 1      | 15    | 0                | TYPE_LIMIT | TIF_GTC | trigger-auction1 |
            | party1 | ETH/DEC19 | buy  | 10     | 9     | 0                | TYPE_LIMIT | TIF_GTC |                  |
            | party2 | ETH/DEC19 | sell | 10     | 9     | 0                | TYPE_LIMIT | TIF_GTC |                  |
    And the parties cancel the following orders:
            | party | reference        |
            | aux   | trigger-auction1 |
            | aux2  | trigger-auction2 |
        Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "ETH/DEC19"
        Given the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value                | time offset |
            | perp.ETH.value | 30000000000000000000 | 0s          |

        ### 6 mins in, still in monitoring auction (fraction outside auction is 5/6, hence the funding payment ends up being 5/6*1000=~833)
        Given the network moves ahead "60" blocks
        Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "ETH/DEC19"
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10800         | 9800          | 833             |
        Given the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value                | time offset |
            | perp.ETH.value | 11000000000000000000 | 0s          |

        # 7 mins in, the auction period will end (fraction outside auction is 5/7, hence the funding payment ends up being 5/7*1000=~714)
        Given the network moves ahead "60" blocks
        Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "ETH/DEC19"
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10800         | 9800          | 714             |
        Then the network moves ahead "1" blocks

        # 8 mins in, still in continuous trading (fraction outside auction is ~6/8, hence the funding payment ends up being 6/8*500=~374)
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10500         | 10000         | 374             |
        Given the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     |
            | party1 | ETH/DEC19 | buy  | 1      | 8     | 0                | TYPE_LIMIT | TIF_GTC |
            | party2 | ETH/DEC19 | sell | 1      | 8     | 1                | TYPE_LIMIT | TIF_GTC |
        And the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value               | time offset |
            | perp.ETH.value | 8000000000000000000 | 0s          |

        # 9 mins in, still in continuous trading (fraction outside auction is ~7/9, hence the funding payment ends up being 7/9*500=~332)
        Given the network moves ahead "60" blocks
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 10142         | 9714          | 332             |
        Given the oracles broadcast data with block time signed with "0xCAFECAFE1":
            | name           | value                | time offset |
            | perp.ETH.value | 14000000000000000000 | 0s          |

        # 10 mins in, still in continuous trading (fraction outside auction is ~8/10, hence the funding payment ends up being 8/10*375=~299)
        Given the network moves ahead "60" blocks
        Then the markets are updated:
            | id        | price monitoring | linear slippage factor | quadratic slippage factor |
            | ETH/DEC19 | default-none     | 0.25                   | 0                         |
        When the parties place the following orders:
            | party  | market id | side | volume | price | resulting trades | type       | tif     |
            | party1 | ETH/DEC19 | buy  | 1      | 30    | 0                | TYPE_LIMIT | TIF_GTC |
            | party2 | ETH/DEC19 | sell | 1      | 30    | 1                | TYPE_LIMIT | TIF_GTC |
        Then the product data for the market "ETH/DEC19" should be:
            | internal twap | external twap | funding payment |
            | 9875          | 10250         | -299            |