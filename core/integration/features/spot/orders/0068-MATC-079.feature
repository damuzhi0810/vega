Feature: Spot market

  Background:
    Given time is updated to "2024-01-01T00:00:00Z"

    Given the following network parameters are set:
      | name                                                | value |
      | network.markPriceUpdateMaximumFrequency             | 0s    |
      | market.value.windowLength                           | 1h    |
    
    Given the following assets are registered:
      | id  | decimal places |
      | ETH | 2              |
      | BTC | 2              |

    Given the fees configuration named "fees-config-1":
      | maker fee | infrastructure fee |
      | 0.01      | 0.03               |
    Given the log normal risk model named "lognormal-risk-model-1":
      | risk aversion | tau  | mu | r   | sigma |
      | 0.001         | 0.01 | 0  | 0.0 | 1.2   |
    And the price monitoring named "price-monitoring-1":
      | horizon | probability | auction extension |
      | 3600    | 0.999       | 10                |

    And the spot markets:
      | id      | name    | base asset | quote asset | risk model             | auction duration | fees          | price monitoring   | decimal places | position decimal places | sla params    |
      | BTC/ETH | BTC/ETH | BTC        | ETH         | lognormal-risk-model-1 | 1                | fees-config-1 | price-monitoring-1 | 2              | 2                       | default-basic |

    # setup accounts
    Given the parties deposit on asset's general account the following amount:
      | party  | asset | amount |
      | party1 | ETH   | 10000  |
      | party1 | BTC   | 1000   |
      | party2 | ETH   | 10000  |
      | party2 | BTC   | 1000   |
      | party5 | ETH   | 10000  |
      | party5 | BTC   | 1000   |
    And the average block duration is "1"

    Scenario: In auction GTC/GTT All MARKET orders are rejected. For product spot: (0068-MATC-079)

      # Place some orders that are invalid in the opening auction
      And the parties place the following orders:
        | party  | market id | side | volume | price | resulting trades | type        | tif     | expires in | reference | error |
        | party1 | BTC/ETH   | sell | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTC |            | sell1     | OrderError: Invalid Persistence |
        | party1 | BTC/ETH   | sell | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTT | 3600       | sell2     | OrderError: Invalid Expiration Datetime |
        | party1 | BTC/ETH   | buy  | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTC |            | buy1      | OrderError: Invalid Persistence |
        | party1 | BTC/ETH   | buy  | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTT | 3600       | buy2      | OrderError: Invalid Expiration Datetime |

      Then the orders should have the following status:
        | party  | reference | status          |
        | party1 | sell1     | STATUS_REJECTED |
        | party1 | sell2     | STATUS_REJECTED |
        | party1 | buy1      | STATUS_REJECTED |
        | party1 | buy2      | STATUS_REJECTED |

    # Place some orders to get out of auction
    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type       | tif     |
      | party1 | BTC/ETH   | buy  | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC |
      | party5 | BTC/ETH   | sell | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC |

    And the opening auction period ends for market "BTC/ETH"
    When the network moves ahead "1" blocks
    Then the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "BTC/ETH"
    And the mark price should be "1000" for the market "BTC/ETH"

    And the market data for the market "BTC/ETH" should be:
      | mark price | trading mode            | horizon | min bound | max bound | target stake | supplied stake | open interest |
      | 1000       | TRADING_MODE_CONTINUOUS | 3600    | 959       | 1042      | 0            | 0              | 0             |

    # Now get into a price monitoring auction
      And the parties place the following orders:
        | party  | market id | side | volume | price | resulting trades | type       | tif     | 
        | party5 | BTC/ETH   | sell | 1      | 1050  | 0                | TYPE_LIMIT | TIF_GTC | 
        | party1 | BTC/ETH   | buy  | 1      | 1050  | 0                | TYPE_LIMIT | TIF_GTC | 

    Then the trading mode should be "TRADING_MODE_MONITORING_AUCTION" for the market "BTC/ETH"

    # Place some orders that are invalid in the price monitoring auction
    And the parties place the following orders:
      | party  | market id | side | volume | price | resulting trades | type        | tif     | expires in | reference | error |
      | party1 | BTC/ETH   | sell | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTC |            | sell3     | OrderError: Invalid Persistence |
      | party1 | BTC/ETH   | sell | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTT | 3600       | sell4     | OrderError: Invalid Expiration Datetime |
      | party1 | BTC/ETH   | buy  | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTC |            | buy3      | OrderError: Invalid Persistence |
      | party1 | BTC/ETH   | buy  | 1      | 1000  | 0                | TYPE_MARKET | TIF_GTT | 3600       | buy4      | OrderError: Invalid Expiration Datetime |

    Then the orders should have the following status:
      | party  | reference | status          |
      | party1 | sell3     | STATUS_REJECTED |
      | party1 | sell4     | STATUS_REJECTED |
      | party1 | buy3      | STATUS_REJECTED |
      | party1 | buy4      | STATUS_REJECTED |