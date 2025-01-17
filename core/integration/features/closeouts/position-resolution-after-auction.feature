Feature: Set up a market with an opening auction, then uncross the book so that the part trading in auction becomes distressed 
  Background:
    Given the markets:
      | id        | quote name | asset | risk model                  | margin calculator         | auction duration | fees         | price monitoring | data source config     | linear slippage factor | quadratic slippage factor | sla params      |
      | ETH/DEC19 | BTC        | BTC   | default-simple-risk-model-4 | default-margin-calculator | 1                | default-none | default-none     | default-eth-for-future | 0.1                    | 0.1                       | default-futures |
    And the parties deposit on asset's general account the following amount:
      | party   | asset | amount    |
      | party1  | BTC   | 20000     |
      | party2a | BTC   | 100000000 |
      | party2b | BTC   | 100000000 |
      | party2c | BTC   | 100000000 |
      | party3  | BTC   | 100000000 |
      | party4  | BTC   | 100000000 |
      | lp      | BTC   | 100000000 |
    And the following network parameters are set:
      | name                                    | value |
      | limits.markets.maxPeggedOrders          | 2     |
      
  @Liquidation
  Scenario:
    When the parties place the following orders:
      | party   | market id | side | volume | price | resulting trades | type       | tif     | reference |
      | party3  | ETH/DEC19 | buy  | 1      | 1000  | 0                | TYPE_LIMIT | TIF_GTC | t3-b-1    |
      | party4  | ETH/DEC19 | sell | 1      | 11000 | 0                | TYPE_LIMIT | TIF_GTC | t4-s-1    |
      | party1  | ETH/DEC19 | buy  | 5      | 10000 | 0                | TYPE_LIMIT | TIF_GFA | t1-b-1    |
      | party2a | ETH/DEC19 | sell | 5      | 10000 | 0                | TYPE_LIMIT | TIF_GFA | t2a-s-1   |
      | party1  | ETH/DEC19 | buy  | 5      | 10000 | 0                | TYPE_LIMIT | TIF_GFA | t1-b-2    |
      | party2b | ETH/DEC19 | sell | 5      | 10001 | 0                | TYPE_LIMIT | TIF_GFA | t2b-s-2   |
      | party1  | ETH/DEC19 | buy  | 4      | 3000  | 0                | TYPE_LIMIT | TIF_GFA | t1-b-3    |
      | party2c | ETH/DEC19 | sell | 3      | 3000  | 0                | TYPE_LIMIT | TIF_GFA | t2c-s-3   |
    And the parties submit the following liquidity provision:
      | id  | party | market id | commitment amount | fee  | lp type    |
      | lp1 | lp    | ETH/DEC19 | 160000            | 0.01 | submission |
      | lp1 | lp    | ETH/DEC19 | 160000            | 0.01 | submission |
    And the parties place the following pegged iceberg orders:
      | party | market id | peak size | minimum visible size | side | pegged reference | volume | offset |
      | lp    | ETH/DEC19 | 27        | 1                    | buy  | MID              | 36     | 100    |
      | lp    | ETH/DEC19 | 27        | 1                    | sell | MID              | 27     | 100    |
    Then the parties should have the following margin levels:
      | party  | market id | maintenance |
      | party1 | ETH/DEC19 | 14000       |
    And the parties should have the following account balances:
      | party  | asset | market id | margin | general |
      | party1 | BTC   | ETH/DEC19 | 16800  |  3200   |
    When the opening auction period ends for market "ETH/DEC19"
    Then the market data for the market "ETH/DEC19" should be:
      | mark price | trading mode            | auction trigger             | extension trigger           | target stake | supplied stake | open interest |
      | 10000      | TRADING_MODE_CONTINUOUS | AUCTION_TRIGGER_UNSPECIFIED | AUCTION_TRIGGER_UNSPECIFIED | 80000        | 160000         | 8             |
    When the network moves ahead "1" blocks
    And debug trades
    And the following trades should be executed:
      | buyer   | price | size | seller  |
      | party1  | 10000 | 3    | party2a |
      | party1  | 10000 | 2    | party2a |
      | party1  | 10000 | 3    | party2c |
      | network | 10000 | 8    | party1  |
      | lp      | 5900  | 8    | network |
    Then the parties should have the following profit and loss:
      | party   | volume | unrealised pnl | realised pnl |
      | party2a | -5     | 0              | 0            |
      | party2c | -3     | 0              | 0            |
      | party1  | 0      | 0              | -20000       |
    And the accumulated liquidity fees should be "472" for the market "ETH/DEC19"
    And the insurance pool balance should be "19528" for the market "ETH/DEC19"
    And the parties should have the following account balances:
      | party  | asset | market id | margin | general  | bond   |
      | party1 | BTC   | ETH/DEC19 | 0      | 0        |        |
      | lp     | BTC   | ETH/DEC19 | 43200  | 99796800 | 160000 |
    # sum of lp accounts = 100019528
    # lp started with 100000000, should've made 8*(10000-5900)=32800 in MTM gains following the closeout,
    # but party1 only had 20000, of which 472 has been put towards liquidity fees, 
    # so only the leftover 19528 was transferred to lp in MTM gains, hence the -13272 realised pnl
