// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/execution (interfaces: OrderBuf)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderBuf is a mock of OrderBuf interface
type MockOrderBuf struct {
	ctrl     *gomock.Controller
	recorder *MockOrderBufMockRecorder
}

// MockOrderBufMockRecorder is the mock recorder for MockOrderBuf
type MockOrderBufMockRecorder struct {
	mock *MockOrderBuf
}

// NewMockOrderBuf creates a new mock instance
func NewMockOrderBuf(ctrl *gomock.Controller) *MockOrderBuf {
	mock := &MockOrderBuf{ctrl: ctrl}
	mock.recorder = &MockOrderBufMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderBuf) EXPECT() *MockOrderBufMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockOrderBuf) Add(arg0 proto.Order) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add
func (mr *MockOrderBufMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockOrderBuf)(nil).Add), arg0)
}

// Flush mocks base method
func (m *MockOrderBuf) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockOrderBufMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockOrderBuf)(nil).Flush))
}