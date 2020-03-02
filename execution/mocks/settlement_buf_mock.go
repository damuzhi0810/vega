// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/execution (interfaces: SettlementBuf)

// Package mocks is a generated GoMock package.
package mocks

import (
	events "code.vegaprotocol.io/vega/events"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSettlementBuf is a mock of SettlementBuf interface
type MockSettlementBuf struct {
	ctrl     *gomock.Controller
	recorder *MockSettlementBufMockRecorder
}

// MockSettlementBufMockRecorder is the mock recorder for MockSettlementBuf
type MockSettlementBufMockRecorder struct {
	mock *MockSettlementBuf
}

// NewMockSettlementBuf creates a new mock instance
func NewMockSettlementBuf(ctrl *gomock.Controller) *MockSettlementBuf {
	mock := &MockSettlementBuf{ctrl: ctrl}
	mock.recorder = &MockSettlementBufMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSettlementBuf) EXPECT() *MockSettlementBufMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockSettlementBuf) Add(arg0 []events.SettlePosition) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add
func (mr *MockSettlementBufMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockSettlementBuf)(nil).Add), arg0)
}

// Flush mocks base method
func (m *MockSettlementBuf) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush
func (mr *MockSettlementBufMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockSettlementBuf)(nil).Flush))
}