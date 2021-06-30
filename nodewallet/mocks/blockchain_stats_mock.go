// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/nodewallet (interfaces: BlockchainStats)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBlockchainStats is a mock of BlockchainStats interface
type MockBlockchainStats struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainStatsMockRecorder
}

// MockBlockchainStatsMockRecorder is the mock recorder for MockBlockchainStats
type MockBlockchainStatsMockRecorder struct {
	mock *MockBlockchainStats
}

// NewMockBlockchainStats creates a new mock instance
func NewMockBlockchainStats(ctrl *gomock.Controller) *MockBlockchainStats {
	mock := &MockBlockchainStats{ctrl: ctrl}
	mock.recorder = &MockBlockchainStatsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchainStats) EXPECT() *MockBlockchainStatsMockRecorder {
	return m.recorder
}

// Height mocks base method
func (m *MockBlockchainStats) Height() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Height")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Height indicates an expected call of Height
func (mr *MockBlockchainStatsMockRecorder) Height() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Height", reflect.TypeOf((*MockBlockchainStats)(nil).Height))
}