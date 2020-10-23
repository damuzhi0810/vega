// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: GovernanceEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	governance "code.vegaprotocol.io/vega/governance"
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockGovernanceEngine is a mock of GovernanceEngine interface
type MockGovernanceEngine struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceEngineMockRecorder
}

// MockGovernanceEngineMockRecorder is the mock recorder for MockGovernanceEngine
type MockGovernanceEngineMockRecorder struct {
	mock *MockGovernanceEngine
}

// NewMockGovernanceEngine creates a new mock instance
func NewMockGovernanceEngine(ctrl *gomock.Controller) *MockGovernanceEngine {
	mock := &MockGovernanceEngine{ctrl: ctrl}
	mock.recorder = &MockGovernanceEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGovernanceEngine) EXPECT() *MockGovernanceEngineMockRecorder {
	return m.recorder
}

// AddVote mocks base method
func (m *MockGovernanceEngine) AddVote(arg0 context.Context, arg1 proto.Vote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVote", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVote indicates an expected call of AddVote
func (mr *MockGovernanceEngineMockRecorder) AddVote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVote", reflect.TypeOf((*MockGovernanceEngine)(nil).AddVote), arg0, arg1)
}

// OnChainTimeUpdate mocks base method
func (m *MockGovernanceEngine) OnChainTimeUpdate(arg0 context.Context, arg1 time.Time) []*governance.ToEnact {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnChainTimeUpdate", arg0, arg1)
	ret0, _ := ret[0].([]*governance.ToEnact)
	return ret0
}

// OnChainTimeUpdate indicates an expected call of OnChainTimeUpdate
func (mr *MockGovernanceEngineMockRecorder) OnChainTimeUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChainTimeUpdate", reflect.TypeOf((*MockGovernanceEngine)(nil).OnChainTimeUpdate), arg0, arg1)
}

// SubmitProposal mocks base method
func (m *MockGovernanceEngine) SubmitProposal(arg0 context.Context, arg1 proto.Proposal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitProposal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitProposal indicates an expected call of SubmitProposal
func (mr *MockGovernanceEngineMockRecorder) SubmitProposal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitProposal", reflect.TypeOf((*MockGovernanceEngine)(nil).SubmitProposal), arg0, arg1)
}