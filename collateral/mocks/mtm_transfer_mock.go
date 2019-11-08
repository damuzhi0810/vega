// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/collateral (interfaces: Transfer)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransfer is a mock of Transfer interface
type MockTransfer struct {
	ctrl     *gomock.Controller
	recorder *MockTransferMockRecorder
}

// MockTransferMockRecorder is the mock recorder for MockTransfer
type MockTransferMockRecorder struct {
	mock *MockTransfer
}

// NewMockTransfer creates a new mock instance
func NewMockTransfer(ctrl *gomock.Controller) *MockTransfer {
	mock := &MockTransfer{ctrl: ctrl}
	mock.recorder = &MockTransferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransfer) EXPECT() *MockTransferMockRecorder {
	return m.recorder
}

// Buy mocks base method
func (m *MockTransfer) Buy() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buy")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Buy indicates an expected call of Buy
func (mr *MockTransferMockRecorder) Buy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buy", reflect.TypeOf((*MockTransfer)(nil).Buy))
}

// ClearPotentials mocks base method
func (m *MockTransfer) ClearPotentials() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearPotentials")
}

// ClearPotentials indicates an expected call of ClearPotentials
func (mr *MockTransferMockRecorder) ClearPotentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPotentials", reflect.TypeOf((*MockTransfer)(nil).ClearPotentials))
}

// Party mocks base method
func (m *MockTransfer) Party() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Party")
	ret0, _ := ret[0].(string)
	return ret0
}

// Party indicates an expected call of Party
func (mr *MockTransferMockRecorder) Party() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Party", reflect.TypeOf((*MockTransfer)(nil).Party))
}

// Price mocks base method
func (m *MockTransfer) Price() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Price")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Price indicates an expected call of Price
func (mr *MockTransferMockRecorder) Price() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Price", reflect.TypeOf((*MockTransfer)(nil).Price))
}

// Sell mocks base method
func (m *MockTransfer) Sell() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sell")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Sell indicates an expected call of Sell
func (mr *MockTransferMockRecorder) Sell() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sell", reflect.TypeOf((*MockTransfer)(nil).Sell))
}

// Size mocks base method
func (m *MockTransfer) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockTransferMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockTransfer)(nil).Size))
}

// Transfer mocks base method
func (m *MockTransfer) Transfer() *proto.Transfer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer")
	ret0, _ := ret[0].(*proto.Transfer)
	return ret0
}

// Transfer indicates an expected call of Transfer
func (mr *MockTransferMockRecorder) Transfer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockTransfer)(nil).Transfer))
}