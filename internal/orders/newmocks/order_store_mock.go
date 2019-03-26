// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/orders (interfaces: OrderStore)

// Package newmocks is a generated GoMock package.
package newmocks

import (
	filtering "code.vegaprotocol.io/vega/internal/filtering"
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderStore is a mock of OrderStore interface
type MockOrderStore struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStoreMockRecorder
}

// MockOrderStoreMockRecorder is the mock recorder for MockOrderStore
type MockOrderStoreMockRecorder struct {
	mock *MockOrderStore
}

// NewMockOrderStore creates a new mock instance
func NewMockOrderStore(ctrl *gomock.Controller) *MockOrderStore {
	mock := &MockOrderStore{ctrl: ctrl}
	mock.recorder = &MockOrderStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderStore) EXPECT() *MockOrderStoreMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockOrderStore) GetByMarket(arg0 context.Context, arg1 string, arg2 *filtering.OrderQueryFilters) ([]*proto.Order, error) {
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockOrderStoreMockRecorder) GetByMarket(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockOrderStore)(nil).GetByMarket), arg0, arg1, arg2)
}

// GetByMarketAndId mocks base method
func (m *MockOrderStore) GetByMarketAndId(arg0 context.Context, arg1, arg2 string) (*proto.Order, error) {
	ret := m.ctrl.Call(m, "GetByMarketAndId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndId indicates an expected call of GetByMarketAndId
func (mr *MockOrderStoreMockRecorder) GetByMarketAndId(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndId", reflect.TypeOf((*MockOrderStore)(nil).GetByMarketAndId), arg0, arg1, arg2)
}

// GetByParty mocks base method
func (m *MockOrderStore) GetByParty(arg0 context.Context, arg1 string, arg2 *filtering.OrderQueryFilters) ([]*proto.Order, error) {
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockOrderStoreMockRecorder) GetByParty(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockOrderStore)(nil).GetByParty), arg0, arg1, arg2)
}

// GetByPartyAndId mocks base method
func (m *MockOrderStore) GetByPartyAndId(arg0 context.Context, arg1, arg2 string) (*proto.Order, error) {
	ret := m.ctrl.Call(m, "GetByPartyAndId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPartyAndId indicates an expected call of GetByPartyAndId
func (mr *MockOrderStoreMockRecorder) GetByPartyAndId(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPartyAndId", reflect.TypeOf((*MockOrderStore)(nil).GetByPartyAndId), arg0, arg1, arg2)
}

// Subscribe mocks base method
func (m *MockOrderStore) Subscribe(arg0 chan<- []proto.Order) uint64 {
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockOrderStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockOrderStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockOrderStore) Unsubscribe(arg0 uint64) error {
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockOrderStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockOrderStore)(nil).Unsubscribe), arg0)
}