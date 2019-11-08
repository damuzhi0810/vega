// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/collateral (interfaces: Accounts)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccounts is a mock of Accounts interface
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsMockRecorder
}

// MockAccountsMockRecorder is the mock recorder for MockAccounts
type MockAccountsMockRecorder struct {
	mock *MockAccounts
}

// NewMockAccounts creates a new mock instance
func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &MockAccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccounts) EXPECT() *MockAccountsMockRecorder {
	return m.recorder
}

// CreateMarketAccounts mocks base method
func (m *MockAccounts) CreateMarketAccounts(arg0 string, arg1 int64) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMarketAccounts", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMarketAccounts indicates an expected call of CreateMarketAccounts
func (mr *MockAccountsMockRecorder) CreateMarketAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMarketAccounts", reflect.TypeOf((*MockAccounts)(nil).CreateMarketAccounts), arg0, arg1)
}

// CreateTraderMarketAccounts mocks base method
func (m *MockAccounts) CreateTraderMarketAccounts(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTraderMarketAccounts", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTraderMarketAccounts indicates an expected call of CreateTraderMarketAccounts
func (mr *MockAccountsMockRecorder) CreateTraderMarketAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTraderMarketAccounts", reflect.TypeOf((*MockAccounts)(nil).CreateTraderMarketAccounts), arg0, arg1)
}

// GetAccountByID mocks base method
func (m *MockAccounts) GetAccountByID(arg0 string) (*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", arg0)
	ret0, _ := ret[0].(*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID
func (mr *MockAccountsMockRecorder) GetAccountByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockAccounts)(nil).GetAccountByID), arg0)
}

// GetMarketAccountsForOwner mocks base method
func (m *MockAccounts) GetMarketAccountsForOwner(arg0, arg1 string) ([]*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketAccountsForOwner", arg0, arg1)
	ret0, _ := ret[0].([]*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketAccountsForOwner indicates an expected call of GetMarketAccountsForOwner
func (mr *MockAccountsMockRecorder) GetMarketAccountsForOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketAccountsForOwner", reflect.TypeOf((*MockAccounts)(nil).GetMarketAccountsForOwner), arg0, arg1)
}

// IncrementBalance mocks base method
func (m *MockAccounts) IncrementBalance(arg0 string, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementBalance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementBalance indicates an expected call of IncrementBalance
func (mr *MockAccountsMockRecorder) IncrementBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementBalance", reflect.TypeOf((*MockAccounts)(nil).IncrementBalance), arg0, arg1)
}

// UpdateBalance mocks base method
func (m *MockAccounts) UpdateBalance(arg0 string, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalance indicates an expected call of UpdateBalance
func (mr *MockAccountsMockRecorder) UpdateBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockAccounts)(nil).UpdateBalance), arg0, arg1)
}