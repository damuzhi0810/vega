Feature: Test market depth events for pegged orders (with BID and ASK price)

  Background:
    Given the initial insurance pool balance is "0" for the markets:
    And the markets:
      | id        | quote name | asset | risk model                  | margin calculator         | auction duration | fees         | price monitoring | oracle config          |
      | ETH/DEC19 | BTC        | BTC   | default-simple-risk-model-2 | default-margin-calculator | 1                | default-none | default-none     | default-eth-for-future |
    And the following network parameters are set:
      | market.auction.minimumDuration |
      | 1                              |
    And the oracles broadcast data signed with "0xDEADBEEF":
      | name             | value |
      | prices.ETH.value | 42    |

  Scenario: Ensure the expect order events for pegged orders are produced for all references
# setup accounts
    Given the traders deposit on asset's general account the following amount:
      | trader           | asset | amount    |
      | sellSideProvider | BTC   | 100000000 |
      | buySideProvider  | BTC   | 100000000 |
      | pegged1          | BTC   | 5000000   |
      | pegged2          | BTC   | 5000000   |
      | pegged3          | BTC   | 5000000   |
      | pegged4          | BTC   | 5000000   |
      | aux              | BTC   | 100000000 |
      | aux2             | BTC   | 100000000 |
# setup pegged orders
    Then the traders place the following pegged orders:
      | trader  | market id | side | volume | reference | offset | price |
      | pegged1 | ETH/DEC19 | sell | 5      | ASK       | 10     | 100   |
      | pegged2 | ETH/DEC19 | sell | 5      | MID       | 15     | 100   |
      | pegged3 | ETH/DEC19 | buy  | 5      | BID       | -10    | 100   |
      | pegged4 | ETH/DEC19 | buy  | 5      | MID       | -10    | 100   |
    Then I see the following order events:
      | trader  | market id | side | volume | reference | offset | price | status        |
      | pegged1 | ETH/DEC19 | sell | 5      | ASK       | 10     | 100   | STATUS_PARKED |
      | pegged2 | ETH/DEC19 | sell | 5      | MID       | 15     | 100   | STATUS_PARKED |
      | pegged3 | ETH/DEC19 | buy  | 5      | BID       | -10    | 100   | STATUS_PARKED |
      | pegged4 | ETH/DEC19 | buy  | 5      | MID       | -10    | 100   | STATUS_PARKED |
# keep things simple: remove the events we've just verified
    And clear order events
    When the traders place the following orders:
      | trader           | market id | side | volume | price | resulting trades | type       | tif     | reference       |
      | sellSideProvider | ETH/DEC19 | sell | 1000   | 120   | 0                | TYPE_LIMIT | TIF_GTC | sell-provider-1 |
      | buySideProvider  | ETH/DEC19 | buy  | 1000   | 80    | 0                | TYPE_LIMIT | TIF_GTC | buy-provider-1  |
      | aux              | ETH/DEC19 | sell | 1      | 100   | 0                | TYPE_LIMIT | TIF_GTC | aux-s-1         |
      | aux2             | ETH/DEC19 | buy  | 1      | 100   | 0                | TYPE_LIMIT | TIF_GTC | aux-b-1         |
    Then I see the following order events:
      | trader           | market id | side | volume | reference | offset | price | status        |
      | sellSideProvider | ETH/DEC19 | sell | 1000   |           | 0      | 120   | STATUS_ACTIVE |
      | buySideProvider  | ETH/DEC19 | buy  | 1000   |           | 0      | 80    | STATUS_ACTIVE |
# Checked out, remove the order events we've checked, now let's have a look at the pegged order events
    And clear order events by reference:
      | trader           | reference       |
      | sellSideProvider | sell-provider-1 |
      | buySideProvider  | buy-provider-1  |
    Then the opening auction period ends for market "ETH/DEC19"
    And the trading mode should be "TRADING_MODE_CONTINUOUS" for the market "ETH/DEC19"
# Now check what happened to our pegged orders
    Then I see the following order events:
      | trader  | market id | side | volume | reference | offset | price | status        |
      | pegged1 | ETH/DEC19 | sell | 5      | ASK       | 10     | 130   | STATUS_ACTIVE |
      | pegged1 | ETH/DEC19 | sell | 5      | ASK       | 10     | 130   | STATUS_ACTIVE |
      | pegged1 | ETH/DEC19 | sell | 5      | ASK       | 10     | 130   | STATUS_ACTIVE |
      | pegged2 | ETH/DEC19 | sell | 5      | MID       | 15     | 115   | STATUS_ACTIVE |
      | pegged3 | ETH/DEC19 | buy  | 5      | BID       | -10    | 70    | STATUS_ACTIVE |
      | pegged4 | ETH/DEC19 | buy  | 5      | MID       | -10    | 90    | STATUS_ACTIVE |