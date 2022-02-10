// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: StakeVerifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "code.vegaprotocol.io/vega/types"
	gomock "github.com/golang/mock/gomock"
)

// MockStakeVerifier is a mock of StakeVerifier interface.
type MockStakeVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockStakeVerifierMockRecorder
}

// MockStakeVerifierMockRecorder is the mock recorder for MockStakeVerifier.
type MockStakeVerifierMockRecorder struct {
	mock *MockStakeVerifier
}

// NewMockStakeVerifier creates a new mock instance.
func NewMockStakeVerifier(ctrl *gomock.Controller) *MockStakeVerifier {
	mock := &MockStakeVerifier{ctrl: ctrl}
	mock.recorder = &MockStakeVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakeVerifier) EXPECT() *MockStakeVerifierMockRecorder {
	return m.recorder
}

// ProcessStakeDeposited mocks base method.
func (m *MockStakeVerifier) ProcessStakeDeposited(arg0 context.Context, arg1 *types.StakeDeposited) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessStakeDeposited", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessStakeDeposited indicates an expected call of ProcessStakeDeposited.
func (mr *MockStakeVerifierMockRecorder) ProcessStakeDeposited(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessStakeDeposited", reflect.TypeOf((*MockStakeVerifier)(nil).ProcessStakeDeposited), arg0, arg1)
}

// ProcessStakeRemoved mocks base method.
func (m *MockStakeVerifier) ProcessStakeRemoved(arg0 context.Context, arg1 *types.StakeRemoved) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessStakeRemoved", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessStakeRemoved indicates an expected call of ProcessStakeRemoved.
func (mr *MockStakeVerifierMockRecorder) ProcessStakeRemoved(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessStakeRemoved", reflect.TypeOf((*MockStakeVerifier)(nil).ProcessStakeRemoved), arg0, arg1)
}