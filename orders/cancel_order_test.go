package orders_test

import (
	"context"
	"fmt"
	"testing"

	"code.vegaprotocol.io/vega/proto"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	cancel = proto.OrderCancellation{
		OrderID:  "order_id",
		MarketID: "market",
		PartyID:  "party",
	}
)

type cancelMatcher struct {
	e proto.Order
}

func TestCancelOrder(t *testing.T) {
	t.Run("Cancel order - success", testCancelOrderSuccess)
	t.Run("Cancel order - order not in storage", testCancelOrderNotFound)
	t.Run("Cancel order - already cancelled", testCancelOrderDuplicate)
	t.Run("Cancel order - order filled", testCancelOrderFilled)
	t.Run("Cancel order - party mismatch", testCancelOrderPartyMismatch)
}

func testCancelOrderSuccess(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx := context.Background()
	arg := cancel
	order := proto.Order{
		Id:        arg.OrderID,
		MarketID:  arg.MarketID,
		PartyID:   arg.PartyID,
		Status:    proto.Order_Active,
		Remaining: 1,
	}

	svc.orderStore.EXPECT().GetByMarketAndID(gomock.Any(), arg.MarketID, arg.OrderID).Times(1).Return(&order, nil)
	svc.block.EXPECT().CancelOrder(gomock.Any(), cancelMatcher{e: order}).Times(1).Return(true, nil)
	pendingOrder, err := svc.svc.CancelOrder(ctx, &arg)
	assert.NotNil(t, pendingOrder)
	assert.NoError(t, err)
}

func testCancelOrderNotFound(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx := context.Background()
	arg := cancel
	osErr := errors.New("orderStore error")

	svc.orderStore.EXPECT().GetByMarketAndID(gomock.Any(), arg.MarketID, arg.OrderID).Times(1).Return(nil, osErr)
	pendingOrder, err := svc.svc.CancelOrder(ctx, &arg)
	assert.Nil(t, pendingOrder)
	assert.Error(t, err)
	assert.Equal(t, osErr, err)
}

func testCancelOrderDuplicate(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx := context.Background()
	arg := cancel
	order := proto.Order{
		Id:        arg.OrderID,
		MarketID:  arg.MarketID,
		PartyID:   arg.PartyID,
		Status:    proto.Order_Cancelled,
		Remaining: 1,
	}

	svc.orderStore.EXPECT().GetByMarketAndID(gomock.Any(), arg.MarketID, arg.OrderID).Times(1).Return(&order, nil)
	pendingOrder, err := svc.svc.CancelOrder(ctx, &arg)
	assert.Nil(t, pendingOrder)
	assert.Error(t, err)
}

func testCancelOrderFilled(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx := context.Background()
	arg := cancel
	order := proto.Order{
		Id:        arg.OrderID,
		MarketID:  arg.MarketID,
		PartyID:   arg.PartyID,
		Status:    proto.Order_Active,
		Remaining: 0,
	}

	svc.orderStore.EXPECT().GetByMarketAndID(gomock.Any(), arg.MarketID, arg.OrderID).Times(1).Return(&order, nil)
	pendingOrder, err := svc.svc.CancelOrder(ctx, &arg)
	assert.Nil(t, pendingOrder)
	assert.Error(t, err)
}

func testCancelOrderPartyMismatch(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx := context.Background()
	arg := cancel
	order := proto.Order{
		Id:        arg.OrderID,
		MarketID:  arg.MarketID,
		PartyID:   fmt.Sprintf("%s-foobar", arg.PartyID),
		Status:    proto.Order_Active,
		Remaining: 1,
	}

	svc.orderStore.EXPECT().GetByMarketAndID(gomock.Any(), arg.MarketID, arg.OrderID).Times(1).Return(&order, nil)
	pendingOrder, err := svc.svc.CancelOrder(ctx, &arg)
	assert.Nil(t, pendingOrder)
	assert.Error(t, err)
}

func (m cancelMatcher) String() string {
	return fmt.Sprintf("%#v", m.e)
}

func (m cancelMatcher) Matches(x interface{}) bool {
	var v proto.Order
	switch val := x.(type) {
	case *proto.Order:
		v = *val
	case proto.Order:
		v = val
	default:
		return false
	}
	return (m.e.Id == v.Id && m.e.MarketID == v.MarketID && m.e.PartyID == v.PartyID && m.e.Status == v.Status && m.e.Remaining == v.Remaining)
}