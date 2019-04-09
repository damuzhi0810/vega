// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/execution (interfaces: MatchingEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMatchingEngine is a mock of MatchingEngine interface
type MockMatchingEngine struct {
	ctrl     *gomock.Controller
	recorder *MockMatchingEngineMockRecorder
}

// MockMatchingEngineMockRecorder is the mock recorder for MockMatchingEngine
type MockMatchingEngineMockRecorder struct {
	mock *MockMatchingEngine
}

// NewMockMatchingEngine creates a new mock instance
func NewMockMatchingEngine(ctrl *gomock.Controller) *MockMatchingEngine {
	mock := &MockMatchingEngine{ctrl: ctrl}
	mock.recorder = &MockMatchingEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMatchingEngine) EXPECT() *MockMatchingEngineMockRecorder {
	return m.recorder
}

// AddOrderBook mocks base method
func (m *MockMatchingEngine) AddOrderBook(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrderBook", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrderBook indicates an expected call of AddOrderBook
func (mr *MockMatchingEngineMockRecorder) AddOrderBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrderBook", reflect.TypeOf((*MockMatchingEngine)(nil).AddOrderBook), arg0)
}

// AmendOrder mocks base method
func (m *MockMatchingEngine) AmendOrder(arg0 *proto.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockMatchingEngineMockRecorder) AmendOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockMatchingEngine)(nil).AmendOrder), arg0)
}

// CancelOrder mocks base method
func (m *MockMatchingEngine) CancelOrder(arg0 *proto.Order) (*proto.OrderCancellationConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0)
	ret0, _ := ret[0].(*proto.OrderCancellationConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockMatchingEngineMockRecorder) CancelOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockMatchingEngine)(nil).CancelOrder), arg0)
}

// RemoveExpiringOrders mocks base method
func (m *MockMatchingEngine) RemoveExpiringOrders(arg0 uint64) []proto.Order {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveExpiringOrders", arg0)
	ret0, _ := ret[0].([]proto.Order)
	return ret0
}

// RemoveExpiringOrders indicates an expected call of RemoveExpiringOrders
func (mr *MockMatchingEngineMockRecorder) RemoveExpiringOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveExpiringOrders", reflect.TypeOf((*MockMatchingEngine)(nil).RemoveExpiringOrders), arg0)
}

// SubmitOrder mocks base method
func (m *MockMatchingEngine) SubmitOrder(arg0 *proto.Order) (*proto.OrderConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOrder", arg0)
	ret0, _ := ret[0].(*proto.OrderConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockMatchingEngineMockRecorder) SubmitOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockMatchingEngine)(nil).SubmitOrder), arg0)
}