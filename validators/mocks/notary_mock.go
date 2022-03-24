// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: Notary)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockNotary is a mock of Notary interface.
type MockNotary struct {
	ctrl     *gomock.Controller
	recorder *MockNotaryMockRecorder
}

// MockNotaryMockRecorder is the mock recorder for MockNotary.
type MockNotaryMockRecorder struct {
	mock *MockNotary
}

// NewMockNotary creates a new mock instance.
func NewMockNotary(ctrl *gomock.Controller) *MockNotary {
	mock := &MockNotary{ctrl: ctrl}
	mock.recorder = &MockNotaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotary) EXPECT() *MockNotaryMockRecorder {
	return m.recorder
}

// IsSigned mocks base method.
func (m *MockNotary) IsSigned(arg0 context.Context, arg1 string, arg2 v1.NodeSignatureKind) ([]v1.NodeSignature, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSigned", arg0, arg1, arg2)
	ret0, _ := ret[0].([]v1.NodeSignature)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// IsSigned indicates an expected call of IsSigned.
func (mr *MockNotaryMockRecorder) IsSigned(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSigned", reflect.TypeOf((*MockNotary)(nil).IsSigned), arg0, arg1, arg2)
}

// OfferSignatures mocks base method.
func (m *MockNotary) OfferSignatures(arg0 v1.NodeSignatureKind, arg1 func(string) []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OfferSignatures", arg0, arg1)
}

// OfferSignatures indicates an expected call of OfferSignatures.
func (mr *MockNotaryMockRecorder) OfferSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferSignatures", reflect.TypeOf((*MockNotary)(nil).OfferSignatures), arg0, arg1)
}

// StartAggregate mocks base method.
func (m *MockNotary) StartAggregate(arg0 string, arg1 v1.NodeSignatureKind, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartAggregate", arg0, arg1, arg2)
}

// StartAggregate indicates an expected call of StartAggregate.
func (mr *MockNotaryMockRecorder) StartAggregate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAggregate", reflect.TypeOf((*MockNotary)(nil).StartAggregate), arg0, arg1, arg2)
}