package risk_test

import (
	"context"
	"fmt"
	"testing"

	"code.vegaprotocol.io/vega/config"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/matching"
	types "code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/risk"
	"code.vegaprotocol.io/vega/risk/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testEngine struct {
	*risk.Engine
	ctrl      *gomock.Controller
	model     *mocks.MockModel
	orderbook *mocks.MockOrderbook
}

// implements the events.Margin interface
type testMargin struct {
	party    string
	size     int64
	buy      int64
	sell     int64
	price    uint64
	transfer *types.Transfer
	asset    string
	margin   uint64
	general  uint64
	market   string
}

var (
	riskResult = types.RiskResult{
		RiskFactors: map[string]*types.RiskFactor{
			"ETH": {
				Market: "ETH/DEC19",
				Short:  .20,
				Long:   .25,
			},
		},
		PredictedNextRiskFactors: map[string]*types.RiskFactor{
			"ETH": {
				Market: "ETH/DEC19",
				Short:  .20,
				Long:   .25,
			},
		},
	}

	riskMinamount      int64  = 250
	riskRequiredMargin int64  = 300
	markPrice          uint64 = 100
)

func TestUpdateMargins(t *testing.T) {
	t.Run("Top up margin test", testMarginTopup)
	t.Run("Noop margin test", testMarginNoop)
	t.Run("Margin too high (overflow)", testMarginOverflow)
	t.Run("Update Margin with orders in book", testMarginWithOrderInBook)
}

func testMarginTopup(t *testing.T) {
	eng := getTestEngine(t, nil)
	defer eng.ctrl.Finish()
	ctx, cfunc := context.WithCancel(context.Background())
	defer cfunc()
	evt := testMargin{
		party:   "trader1",
		size:    1,
		price:   1000,
		asset:   "ETH",
		margin:  10,     // required margin will be > 30 so ensure we don't have enough
		general: 100000, // plenty of balance for the transfer anyway
		market:  "ETH/DEC19",
	}
	eng.orderbook.EXPECT().GetCloseoutPrice(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(volume uint64, side types.Side) (uint64, error) {
			return markPrice, nil
		})
	evts := []events.Margin{evt}
	resp := eng.UpdateMarginsOnSettlement(ctx, evts, markPrice)
	assert.Equal(t, 1, len(resp))
	// ensure we get the correct transfer request back, correct amount etc...
	trans := resp[0].Transfer()
	assert.Equal(t, int64(20), trans.Amount.Amount)
	// min = 15 so we go back to search level
	assert.Equal(t, int64(15), trans.Amount.MinAmount)
	assert.Equal(t, types.TransferType_MARGIN_LOW, trans.Type)
}

func testMarginNoop(t *testing.T) {
	eng := getTestEngine(t, nil)
	defer eng.ctrl.Finish()
	ctx, cfunc := context.WithCancel(context.Background())
	defer cfunc()
	evt := testMargin{
		party:   "trader1",
		size:    1,
		price:   1000,
		asset:   "ETH",
		margin:  30,     // more than enough margin to cover the position, not enough to trigger transfer to general
		general: 100000, // plenty of balance for the transfer anyway
		market:  "ETH/DEC19",
	}
	eng.orderbook.EXPECT().GetCloseoutPrice(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(volume uint64, side types.Side) (uint64, error) {
			return markPrice, nil
		})

	evts := []events.Margin{evt}
	resp := eng.UpdateMarginsOnSettlement(ctx, evts, markPrice)
	assert.Equal(t, 0, len(resp))
}

func testMarginOverflow(t *testing.T) {
	eng := getTestEngine(t, nil)
	defer eng.ctrl.Finish()
	ctx, cfunc := context.WithCancel(context.Background())
	defer cfunc()
	evt := testMargin{
		party:   "trader1",
		size:    1,
		price:   1000,
		asset:   "ETH",
		margin:  500,    // required margin will be > 35 (release), so ensure we don't have enough
		general: 100000, // plenty of balance for the transfer anyway
		market:  "ETH/DEC19",
	}
	eng.orderbook.EXPECT().GetCloseoutPrice(gomock.Any(), gomock.Any()).Times(1).
		DoAndReturn(func(volume uint64, side types.Side) (uint64, error) {
			return markPrice, nil
		})
	evts := []events.Margin{evt}
	resp := eng.UpdateMarginsOnSettlement(ctx, evts, markPrice)
	assert.Equal(t, 1, len(resp))

	// ensure we get the correct transfer request back, correct amount etc...
	trans := resp[0].Transfer()
	assert.Equal(t, int64(465), trans.Amount.Amount)
	// assert.Equal(t, riskMinamount-int64(evt.margin), trans.Amount.MinAmount)
	assert.Equal(t, types.TransferType_MARGIN_HIGH, trans.Type)
}

// implementation of the test from the specs
// https://gitlab.com/vega-protocol/product/blob/master/specs/0019-margin-calculator.md#pseudo-code-examples
func testMarginWithOrderInBook(t *testing.T) {
	// custom risk factors
	r := &types.RiskResult{
		RiskFactors: map[string]*types.RiskFactor{
			"ETH": {
				Market: "ETH/DEC19",
				Short:  .11,
				Long:   .10,
			},
		},
		PredictedNextRiskFactors: map[string]*types.RiskFactor{
			"ETH": {
				Market: "ETH/DEC19",
				Short:  .11,
				Long:   .10,
			},
		},
	}
	// custom scaling factor
	mc := &types.MarginCalculator{
		ScalingFactors: &types.ScalingFactors{
			SearchLevel:       1.1,
			InitialMargin:     1.2,
			CollateralRelease: 1.3,
		},
	}

	var markPrice int64 = 144

	// list of order in the book before the test happen
	ordersInBook := []struct {
		volume int64
		price  int64
		tid    string
		side   types.Side
	}{
		// asks
		// {volume: 3, price: 258, tid: "t1", side: types.Side_Sell},
		// {volume: 5, price: 240, tid: "t2", side: types.Side_Sell},
		// {volume: 3, price: 188, tid: "t3", side: types.Side_Sell},
		// bids

		{volume: 1, price: 120, tid: "t4", side: types.Side_Buy},
		{volume: 4, price: 240, tid: "t5", side: types.Side_Buy},
		{volume: 7, price: 258, tid: "t6", side: types.Side_Buy},
	}

	marketID := "testingmarket"

	conf := config.NewDefaultConfig("")
	log := logging.NewTestLogger()
	ctrl := gomock.NewController(t)
	model := mocks.NewMockModel(ctrl)

	// instanciate the book then fil it with the orders

	book := matching.NewOrderBook(
		log, conf.Matching, marketID, uint64(markPrice), false)

	for _, v := range ordersInBook {
		o := &types.Order{
			Id:          fmt.Sprintf("o-%v-%v", v.tid, marketID),
			MarketID:    marketID,
			Side:        v.side,
			Price:       uint64(v.price),
			Size:        uint64(v.volume),
			Remaining:   uint64(v.volume),
			TimeInForce: types.Order_GTT,
			Type:        types.Order_LIMIT,
			Status:      types.Order_Active,
			ExpiresAt:   10000,
		}
		_, err := book.SubmitOrder(o)
		assert.Nil(t, err)
	}

	testE := risk.NewEngine(log, conf.Risk, mc, model, r, book)
	evt := testMargin{
		party:   "tx",
		size:    10,
		buy:     4,
		sell:    8,
		price:   144,
		asset:   "ETH",
		margin:  500,
		general: 100000,
		market:  "ETH/DEC19",
	}
	riskevt := testE.UpdateMarginOnNewOrder(evt, uint64(markPrice))
	assert.NotNil(t, riskevt)
	if riskevt == nil {
		t.Fatal("expecting non nil risk update")
	}
	margins := riskevt.MarginLevels()
	assert.Equal(t, int64(1131), margins.MaintenanceMargin)
	assert.Equal(t, int64(1131*mc.ScalingFactors.SearchLevel), margins.SearchLevel)
	assert.Equal(t, int64(1131*mc.ScalingFactors.InitialMargin), margins.InitialMargin)
	assert.Equal(t, int64(1131*mc.ScalingFactors.CollateralRelease), margins.CollateralReleaseLevel)
}

func getTestEngine(t *testing.T, initialRisk *types.RiskResult) *testEngine {
	if initialRisk == nil {
		cpy := riskResult
		initialRisk = &cpy // this is just a shallow copy, so might be worth creating a deep copy depending on the test
	}
	ctrl := gomock.NewController(t)
	model := mocks.NewMockModel(ctrl)
	conf := risk.NewDefaultConfig()
	ob := mocks.NewMockOrderbook(ctrl)
	engine := risk.NewEngine(
		logging.NewTestLogger(),
		conf,
		getMarginCalculator(),
		model,
		initialRisk,
		ob,
	)
	return &testEngine{
		Engine:    engine,
		ctrl:      ctrl,
		model:     model,
		orderbook: ob,
	}
}

func getMarginCalculator() *types.MarginCalculator {
	return &types.MarginCalculator{
		ScalingFactors: &types.ScalingFactors{
			SearchLevel:       1.1,
			InitialMargin:     1.2,
			CollateralRelease: 1.4,
		},
	}
}

func (m testMargin) Party() string {
	return m.party
}

func (m testMargin) MarketID() string {
	return m.market
}

func (m testMargin) Asset() string {
	return m.asset
}

func (m testMargin) MarginBalance() uint64 {
	return m.margin
}

func (m testMargin) GeneralBalance() uint64 {
	return m.general
}

func (m testMargin) Price() uint64 {
	return m.price
}

func (m testMargin) Buy() int64 {
	return m.buy
}

func (m testMargin) Sell() int64 {
	return m.sell
}

func (m testMargin) Size() int64 {
	return m.size
}

func (m testMargin) ClearPotentials() {}

func (m testMargin) Transfer() *types.Transfer {
	return m.transfer
}