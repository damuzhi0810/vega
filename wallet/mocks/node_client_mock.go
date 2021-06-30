// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet (interfaces: NodeClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNodeClient is a mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// LastBlockHeight mocks base method
func (m *MockNodeClient) LastBlockHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlockHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastBlockHeight indicates an expected call of LastBlockHeight
func (mr *MockNodeClientMockRecorder) LastBlockHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockHeight", reflect.TypeOf((*MockNodeClient)(nil).LastBlockHeight), arg0)
}