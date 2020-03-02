package matching

import (
	"testing"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/stretchr/testify/assert"
)

func TestOrderBookSimple_simpleLimitBuy(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(100), price)
	assert.Equal(t, uint64(1), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 1)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(1))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 1)
}

func TestOrderBookSimple_simpleLimitSell(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestOfferPriceAndVolume()
	assert.Equal(t, uint64(100), price)
	assert.Equal(t, uint64(1), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 1)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(1))
	assert.Equal(t, len(book.ordersByID), 1)
}

func TestOrderBookSimple_simpleMarketBuy(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_IOC,
		Type:        types.Order_MARKET,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleMarketSell(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_IOC,
		Type:        types.Order_MARKET,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

/*
 * NETWORK orders are the same as MARKET+FOK order so should not stay on the book
 * Make sure orders are cancelled and the book is left empty
 */
func TestOrderBookSimple_simpleNetworkBuy(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_FOK,
		Type:        types.Order_NETWORK,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleNetworkSell(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        1,
		Remaining:   1,
		TimeInForce: types.Order_FOK,
		Type:        types.Order_NETWORK,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

/*
 * Now we test simple orders against a book with orders in
 */
func TestOrderBookSimple_simpleLimitBuyFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleLimitSellFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleMarketBuyFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_IOC,
		Type:        types.Order_MARKET,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleMarketSellFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_IOC,
		Type:        types.Order_MARKET,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleNetworkBuyFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_FOK,
		Type:        types.Order_NETWORK,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_simpleNetworkSellFill(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "B",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_FOK,
		Type:        types.Order_NETWORK,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

func TestOrderBookSimple_FillAgainstGTTOrder(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTT,
		Type:        types.Order_LIMIT,
		ExpiresAt:   10,
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_FOK,
		Type:        types.Order_NETWORK,
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(confirm.Trades))

	price, volume := book.BestBidPriceAndVolume()
	assert.Equal(t, uint64(0), price)
	assert.Equal(t, uint64(0), volume)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}

type MarketPos struct {
	size, buy, sell int64
	party           string
	price           uint64
}

func (m MarketPos) Party() string {
	return m.party
}

func (m MarketPos) Size() int64 {
	return m.size
}

func (m MarketPos) Buy() int64 {
	return m.buy
}

func (m MarketPos) Sell() int64 {
	return m.sell
}

func (m MarketPos) Price() uint64 {
	return m.price
}

func TestOrderBookSimple_CancelDistressedOrders(t *testing.T) {
	market := "testMarket"
	book := getTestOrderBook(t, market)
	defer book.Finish()
	order := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Buy,
		Price:       100,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTT,
		Type:        types.Order_LIMIT,
		ExpiresAt:   10,
		Id:          "v0000000000000-0000001",
	}
	confirm, err := book.SubmitOrder(&order)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	order2 := types.Order{
		MarketID:    market,
		PartyID:     "A",
		Side:        types.Side_Sell,
		Price:       101,
		Size:        10,
		Remaining:   10,
		TimeInForce: types.Order_GTC,
		Type:        types.Order_LIMIT,
		Id:          "v0000000000000-0000002",
	}
	confirm, err = book.SubmitOrder(&order2)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(confirm.Trades))

	assert.Equal(t, book.getNumberOfBuyLevels(), 1)
	assert.Equal(t, book.getNumberOfSellLevels(), 1)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(10))
	assert.Equal(t, book.getTotalSellVolume(), uint64(10))
	assert.Equal(t, len(book.ordersByID), 2)

	// Now create a structure to contain the details of distressed party "A" and send them to be cancelled.
	parties := []events.MarketPosition{
		MarketPos{
			party: "A",
		},
	}
	orders, err := book.RemoveDistressedOrders(parties)
	assert.Equal(t, len(orders), 2)
	assert.Equal(t, book.getNumberOfBuyLevels(), 0)
	assert.Equal(t, book.getNumberOfSellLevels(), 0)
	assert.Equal(t, book.getTotalBuyVolume(), uint64(0))
	assert.Equal(t, book.getTotalSellVolume(), uint64(0))
	assert.Equal(t, len(book.ordersByID), 0)
}