package datastore

import (
	"testing"
	"vega/proto"

	"github.com/stretchr/testify/assert"
)

const testMarket = "market"

func TestNewMemStore_ReturnsNewMemStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	assert.NotNil(t, memStore)
}

func TestNewMemStore_ReturnsNewTradeStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)
	assert.NotNil(t, newTradeStore)
}

func TestNewMemStore_ReturnsNewOrderStoreInstance(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	assert.NotNil(t, newOrderStore)
}

func TestMemStore_PostAndGetNewOrder(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)
}

func TestMemStore_PostDuplicateOrder(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)
	err = newOrderStore.Post(order)
	assert.Error(t, err, "order exists in store")
}

func TestMemStore_PostOrderToNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: "GBP/EUR19",
		},
	}
	err := newOrderStore.Post(order)
	assert.Error(t, err, "market does not exist")
}

func TestMemStore_PostPutAndGetExistingOrder(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		Order: msg.Order{
			Id:     "c471bdd5f381aa3654d98f4591eaa968",
			Market: testMarket,
			Party:  "tester",
			Price:  100,
			Size:   1,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, uint64(100), o.Price)
	assert.Equal(t, uint64(1), o.Size)

	order.Price = 1000
	order.Size = 5

	err = newOrderStore.Put( order)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)
	assert.Equal(t, uint64(1000), o.Price)
	assert.Equal(t, uint64(5), o.Size)
}


func TestMemStore_PutNoneExistentOrder(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},
	}
	err := newOrderStore.Put(order)
	assert.Error(t, err, "order not found in store")
}

func TestMemStore_PutOrderToNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: "GBP/EUR19",
		},
	}
	err := newOrderStore.Put(order)
	assert.Error(t, err, "market does not exist")
}

func TestMemStore_PostAndDeleteOrder(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)

	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: testMarket,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	o, err := newOrderStore.Get(testMarket, order.Id)
	assert.Nil(t, err)
	assert.Equal(t, order, o)

	err = newOrderStore.Delete(o)
	assert.Nil(t, err)

	o, err = newOrderStore.Get(testMarket, order.Id)
	assert.Error(t, err)
}

func TestMemStore_DeleteOrderFromNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var order = Order{
		Order: msg.Order{
			Id:     "45305210ff7a9bb9450b1833cc10368a",
			Market: "GBP/EUR19",
		},
	}
	err := newOrderStore.Delete(order)
	assert.Error(t, err, "market does not exist")
}


func TestMemStore_GetOrderForNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	_, err := newOrderStore.Get("UNKNOWN", "ID")
	assert.Error(t, err, "market does not exist")
}

func TestMemStore_PostAndGetTrade(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	var trade = Trade{
		Trade:   msg.Trade{Market: testMarket},
		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
	}

	var order = Order{
		Order: msg.Order{
			Id:     "d41d8cd98f00b204e9800998ecf8427e",
			Market: testMarket,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	err = newTradeStore.Post(trade)
	assert.Nil(t, err)

	tr, err := newTradeStore.Get(testMarket, trade.Id)
	assert.Equal(t, trade, tr)
}

func TestMemStore_PutAndDeleteTrade(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	var order = Order{
		Order: msg.Order{Id: "d41d8cd98f00b204e9800998ecf8427e", Market: testMarket},
	}
	var trade = Trade{
		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
		Trade:   msg.Trade{Market: testMarket},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	err = newTradeStore.Post(trade)
	assert.Nil(t, err)

	tr, err := newTradeStore.Get(testMarket, trade.Id)
	assert.Equal(t, trade, tr)

	err = newTradeStore.Delete(tr)
	assert.Nil(t, err)

	tr, err = newTradeStore.Get(testMarket, trade.Id)
	assert.Error(t, err)
}

func TestMemStore_PostTradeOrderNotFound(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)
	trade := Trade{
		Trade: msg.Trade{
			Id:     "one",
			Market: testMarket,
		},
		OrderId: "mystery",
	}
	err := newTradeStore.Post(trade)
	assert.Error(t, err)
}

func TestMemStore_PostAndFindByOrderId(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	trade1 := Trade{
		Trade: msg.Trade{
			Id:     "one",
			Market: testMarket,
		},
		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
	}
	trade2 := Trade{
		Trade: msg.Trade{
			Id:     "two",
			Market: testMarket,
		},
		OrderId: "d41d8cd98f00b204e9800998ecf8427e",
	}
	order := Order{
		Order: msg.Order{
			Id:     "d41d8cd98f00b204e9800998ecf8427e",
			Market: testMarket,
		},
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)

	err = newTradeStore.Post(trade1)
	assert.Nil(t, err)

	err = newTradeStore.Post(trade2)
	assert.Nil(t, err)

	trades, err := newTradeStore.GetByOrderId(testMarket, order.Id, GetParams{Limit: 12345} )
	assert.Nil(t, err)

	assert.Equal(t, 2, len(trades))
	assert.Equal(t, "one", trades[0].Id)
	assert.Equal(t, "two", trades[1].Id)
}

func TestMemStore_GetAllOrdersForMarket(t *testing.T) {

	var tests = []struct {
		inMarkets  []string
		inOrders  []Order
		inParty   string
		inLimit   uint64
		inMarket string
		outOrdersCount int
	}{
		{
			inMarkets: []string { testMarket, "marketZ" },
			inOrders: []Order {
				{
					Order: msg.Order{
						Id:     "d41d8cd98f00b204e9800998ecf8427e",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "ad2dc275947362c45893bbeb30fc3098",
						Market: "marketZ",
						Party: "party",
					},
				},
				{
					Order: msg.Order{
						Id:     "4e8e41367997cfe705d62ea80592cbcc",
						Market: testMarket,
						Party: "party",
					},
				},
			},
			inParty: "partyA",
			inLimit: 5000,
			inMarket: testMarket,
			outOrdersCount: 1,
		},
		{
			inMarkets: []string { testMarket, "marketZ" },
			inOrders: []Order {
				{
					Order: msg.Order{
						Id:     "d41d8cd98f00b204e9800998ecf8427e",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "ad2dc275947362c45893bbeb30fc3098",
						Market: "marketZ",
						Party: "party",
					},
				},
				{
					Order: msg.Order{
						Id:     "4e8e41367997cfe705d62ea80592cbcc",
						Market: testMarket,
						Party: "party",
					},
				},
			},
			inParty: "",
			inLimit: 5000,
			inMarket: testMarket,
			outOrdersCount: 2,
		},
		{
			inMarkets: []string { testMarket },
			inOrders: []Order {
				{
					Order: msg.Order{
						Id:     "d41d8cd98f00b204e9800998ecf8427e",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "ad2dc275947362c45893bbeb30fc3098",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "4e8e41367997cfe705d62ea80592cbcc",
						Market: testMarket,
						Party: "partyB",
					},
				},
			},
			inLimit: 2,
			inParty: "partyA",
			inMarket: testMarket,
			outOrdersCount: 2,
		},
		{
			inMarkets: []string { testMarket, "marketY", "marketZ" },
			inOrders: []Order {
				{
					Order: msg.Order{
						Id:     "d41d8cd98f00b204e9800998ecf8427e",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "ad2dc275947362c45893bbeb30fc3098",
						Market: "marketY",
						Party: "partyB",
					},
				},
				{
					Order: msg.Order{
						Id:     "4e8e41367997cfe705d62ea80592cbcc",
						Market: "marketZ",
						Party: "partyB",
					},
				},
			},
			inParty: "",
			inLimit: 9999,
			inMarket: "",
			outOrdersCount: 3,
		},
		{
			inMarkets: []string { testMarket, "marketY", "marketZ" },
			inOrders: []Order {
				{
					Order: msg.Order{
						Id:     "d41d8cd98f00b204e9800998ecf8427e",
						Market: testMarket,
						Party: "partyA",
					},
				},
				{
					Order: msg.Order{
						Id:     "ad2dc275947362c45893bbeb30fc3098",
						Market: "marketY",
						Party: "partyB",
					},
				},
				{
					Order: msg.Order{
						Id:     "4e8e41367997cfe705d62ea80592cbcc",
						Market: "marketZ",
						Party: "partyB",
					},
				},
			},
			inParty: "partyB",
			inLimit: 9999,
			inMarket: "",
			outOrdersCount: 2,
		},
	}
	for _, tt := range tests {
		var memStore = NewMemStore(tt.inMarkets)
		var newOrderStore = NewOrderStore(&memStore)

		for _, order := range tt.inOrders {
			err := newOrderStore.Post(order)
			assert.Nil(t, err)
		}

		params := GetParams{Limit: tt.inLimit}
		orders, err := newOrderStore.GetAll(tt.inMarket, tt.inParty, params)
		assert.Nil(t, err)
		assert.Equal(t, tt.outOrdersCount, len(orders))
	}
}

func TestMemStore_GetAllOrdersForNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	o, err := newOrderStore.GetAll("UNKNOWN", "", GetParams{ Limit: GetParamsLimitDefault })
	assert.Error(t, err, "market does not exist")
	assert.Nil(t, o)
}

func TestMemStore_GetAllTradesForMarket(t *testing.T) {
	otherMarket := "another"
	var memStore = NewMemStore([]string{testMarket, otherMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	orderIdA := "d41d8cd98f00b204e9800998ecf8427e"

	orderA := Order{
		Order: msg.Order{
			Id:     orderIdA,
			Market: testMarket,
		},
	}

	tradeA := Trade{
		Trade: msg.Trade{
			Id: "c0e8490aa4b1d0071ae8f01cdf45c6aa",
			Price: 1000,
			Market: testMarket,
		},
		OrderId: orderIdA,
	}
	tradeB := Trade{
		Trade: msg.Trade{
			Id: "d41d8cd98fsb204e9800998ecf8427e",
			Price: 2000,
			Market: testMarket,
		},
		OrderId: orderIdA,
	}

	err := newOrderStore.Post(orderA)
	assert.Nil(t, err)
	err = newTradeStore.Post(tradeA)
	assert.Nil(t, err)
	err = newTradeStore.Post(tradeB)
	assert.Nil(t, err)

	trades, err := newTradeStore.GetAll(testMarket, GetParams{Limit: 12345})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(trades))
}

func TestMemStore_GetAllTradesForNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)
	o, err := newTradeStore.GetAll("UNKNOWN", GetParams{ Limit: GetParamsLimitDefault })
	assert.Error(t, err, "market does not exist")
	assert.Nil(t, o)
}

func TestMemStore_PutTrade(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newOrderStore = NewOrderStore(&memStore)
	var newTradeStore = NewTradeStore(&memStore)

	orderId := "d41d8cd98f00b204e9800998ecf8427e"
	order := Order{
		Order: msg.Order{
			Id:     orderId,
			Market: testMarket,
		},
	}

	tradeId := "c0e8490aa4b1d0071ae8f01cdf45c6aa"
	trade := Trade{
		Trade: msg.Trade{
			Id: tradeId,
			Price: 1000,
			Market: testMarket,
		},
		OrderId: orderId,
	}

	err := newOrderStore.Post(order)
	assert.Nil(t, err)
	err = newTradeStore.Post(trade)
	assert.Nil(t, err)


	tradeOut, err := newTradeStore.Get(testMarket, tradeId)
	assert.Nil(t, err)
	assert.Equal(t, uint64(1000), tradeOut.Price)

	trade = Trade{
		Trade: msg.Trade{
			Id: tradeId,
			Price: 9000,
			Market: testMarket,
		},
		OrderId: orderId,
	}

	err = newTradeStore.Put(trade)
	assert.Nil(t, err)

	tradeOut, err = newTradeStore.Get(testMarket, tradeId)
	assert.Nil(t, err)
	assert.Equal(t, uint64(9000), tradeOut.Price)
}

func TestMemStore_PutGetAndDeleteTradeForNoneExistentMarket(t *testing.T) {
	var memStore = NewMemStore([]string{testMarket})
	var newTradeStore = NewTradeStore(&memStore)

	trade := Trade{
		Trade: msg.Trade{
			Id: "A",
			Price: 9000,
			Market: "UNKNOWN",
		},
		OrderId: "Z",
	}

	err := newTradeStore.Put(trade)
	assert.Error(t, err, "market does not exist")
	
	_, err = newTradeStore.Get("UNKNOWN", "ID")
	assert.Error(t, err, "market does not exist")

	err = newTradeStore.Delete(trade)
	assert.Error(t, err, "market does not exist")

}

func TestMemOrder_ToString(t *testing.T) {
	orderId := "d41d8cd98f00b204e9800998ecf8427e"
	order := Order{
		Order: msg.Order{
			Id:     orderId,
			Market: testMarket,
		},
	}
	memOrder := memOrder{
		order: order,
	}
	assert.Equal(t, "memOrder::order-id=d41d8cd98f00b204e9800998ecf8427e", memOrder.String())
}

func TestMemTrade_ToString(t *testing.T) {
	tradeId := "c0e8490aa4b1d0071ae8f01cdf45c6aa"
	orderId := "d41d8cd98f00b204e9800998ecf8427e"
	trade := Trade{
		Trade: msg.Trade{
			Id: tradeId,
			Price: 9000,
			Market: testMarket,
		},
		OrderId: orderId,
	}
	memTrade := memTrade{
		trade: trade,
	}
	assert.Equal(t, "memTrade::trade-id=c0e8490aa4b1d0071ae8f01cdf45c6aa", memTrade.String())
}