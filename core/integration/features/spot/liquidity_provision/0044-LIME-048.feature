Feature: Spot market SLA

  Scenario: 001 0044-LIME-048, check SLA when fraction_of_time_on_book < market.liquidity.commitmentMinTimeFraction

    Given time is updated to "2023-07-20T00:00:00Z"

    Given the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0         | 0                  |
    Given the log normal risk model named "lognormal-risk-model-1":
      | risk aversion | tau  | mu | r   | sigma |
      | 0.001         | 0.01 | 0  | 0.0 | 1.2   |

    And the price monitoring named "price-monitoring-1":
      | horizon | probability | auction extension |
      | 360000  | 0.999       | 300               |

    And the liquidity sla params named "SLA-1":
      | price range | commitment min time fraction | performance hysteresis epochs | sla competition factor |
      | 1           | 0.6                          | 2                             | 0.2                    |

    Given the following assets are registered:
      | id  | decimal places |
      | ETH | 1              |
      | BTC | 1              |

    And the following network parameters are set:
      | name                                                | value |
      | network.markPriceUpdateMaximumFrequency             | 2s    |
      | market.liquidity.earlyExitPenalty                   | 0.25  |
      | market.liquidity.bondPenaltyParameter               | 0.2   |
      | market.liquidity.sla.nonPerformanceBondPenaltySlope | 0.7   |
      | market.liquidity.sla.nonPerformanceBondPenaltyMax   | 0.6   |
      | market.liquidity.maximumLiquidityFeeFactorLevel     | 0.4   |
      | validators.epoch.length                             | 10s   |

    And the spot markets:
      | id      | name    | base asset | quote asset | risk model             | auction duration | fees          | price monitoring   | sla params |
      | BTC/ETH | BTC/ETH | BTC        | ETH         | lognormal-risk-model-1 | 1                | fees-config-1 | price-monitoring-1 | SLA-1      |
    And the following network parameters are set:
      | name                                             | value |
      | market.liquidity.providersFeeCalculationTimeStep | 1s    |
      | market.liquidity.stakeToCcyVolume                | 1     |

    Given the parties deposit on asset's general account the following amount:
      | party  | asset | amount |
      | party1 | ETH   | 10000  |
      | party2 | BTC   | 500    |
      | lp1    | ETH   | 40000  |
      | lp1    | BTC   | 2000   |
      | lp2    | ETH   | 4000   |
      | lp2    | BTC   | 60     |

    And the average block duration is "1"

    Given the liquidity monitoring parameters:
      | name               | triggering ratio | time window | scaling factor |
      | updated-lqm-params | 0.2              | 20s         | 0.8            |

    When the spot markets are updated:
      | id      | liquidity monitoring | linear slippage factor | quadratic slippage factor |
      | BTC/ETH | updated-lqm-params   | 0.5                    | 0.5                       |

    When the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee | lp type    |
      | lp1 | lp1   | BTC/ETH   | 3000              | 0.1 | submission |

    Then the network moves ahead "4" blocks
    And the network treasury balance should be "0" for the asset "ETH"
    And the global insurance pool balance should be "0" for the asset "ETH"
    And the global insurance pool balance should be "0" for the asset "BTC"
    And the party "lp1" lp liquidity fee account balance should be "0" for the market "BTC/ETH"
    Then the party "lp1" lp liquidity bond account balance should be "3000" for the market "BTC/ETH"

    # place orders and generate trades
    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     | reference    |
      | lp1    | BTC/ETH   | buy  | 600    | 5     | 0                | TYPE_LIMIT | TIF_GTC | lp1-b        |
      | party1 | BTC/ETH   | buy  | 6      | 8     | 0                | TYPE_LIMIT | TIF_GTC | party-order5 |
      | party1 | BTC/ETH   | buy  | 1      | 15    | 0                | TYPE_LIMIT | TIF_GTC | party-order3 |
      | party2 | BTC/ETH   | sell | 1      | 15    | 0                | TYPE_LIMIT | TIF_GTC | party-order4 |
      | party2 | BTC/ETH   | sell | 6      | 24    | 0                | TYPE_LIMIT | TIF_GTC | party-order6 |
      | lp1    | BTC/ETH   | sell | 120    | 25    | 0                | TYPE_LIMIT | TIF_GTC | lp1-s        |

    Then the network moves ahead "1" blocks

    Then the following trades should be executed:
      | buyer  | price | size | seller |
      | party1 | 15    | 1    | party2 |

    Then "lp1" should have holding account balance of "30000" for asset "ETH"
    Then "lp1" should have general account balance of "7000" for asset "ETH"
    Then the party "lp1" lp liquidity bond account balance should be "3000" for the market "BTC/ETH"
    Then "lp1" should have holding account balance of "1200" for asset "BTC"
    Then "lp1" should have general account balance of "800" for asset "BTC"

    When the network moves ahead "2" blocks
    Then the parties cancel the following orders:
      | party | reference |
      | lp1   | lp1-b     |
      | lp1   | lp1-s     |

    #0011-MARA-020 cancelling an order releases the holding amount back to user's general account.
    Then "lp1" should have holding account balance of "0" for asset "ETH"
    Then "lp1" should have general account balance of "37000" for asset "ETH"
    Then "lp1" should have holding account balance of "0" for asset "BTC"
    Then "lp1" should have general account balance of "2000" for asset "BTC"

    Then the market data for the market "BTC/ETH" should be:
      | mark price | trading mode            | auction trigger             | target stake | supplied stake | open interest |
      | 15         | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | 2400         | 3000           | 0             |

    Then the network moves ahead "7" blocks
    And the network treasury balance should be "1322" for the asset "ETH"
    #fraction_of_time_on_book=2/9 = 0.222
    #0.7 * (1 - 0.222 / 0.6)*3000 = 1322
    Then the party "lp1" lp liquidity bond account balance should be "1678" for the market "BTC/ETH"



