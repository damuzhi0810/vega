// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: OracleService)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/vega/proto/oracles/v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOracleService is a mock of OracleService interface
type MockOracleService struct {
	ctrl     *gomock.Controller
	recorder *MockOracleServiceMockRecorder
}

// MockOracleServiceMockRecorder is the mock recorder for MockOracleService
type MockOracleServiceMockRecorder struct {
	mock *MockOracleService
}

// NewMockOracleService creates a new mock instance
func NewMockOracleService(ctrl *gomock.Controller) *MockOracleService {
	mock := &MockOracleService{ctrl: ctrl}
	mock.recorder = &MockOracleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOracleService) EXPECT() *MockOracleServiceMockRecorder {
	return m.recorder
}

// GetOracleDataBySpecID mocks base method
func (m *MockOracleService) GetOracleDataBySpecID(arg0 string) ([]v1.OracleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOracleDataBySpecID", arg0)
	ret0, _ := ret[0].([]v1.OracleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOracleDataBySpecID indicates an expected call of GetOracleDataBySpecID
func (mr *MockOracleServiceMockRecorder) GetOracleDataBySpecID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOracleDataBySpecID", reflect.TypeOf((*MockOracleService)(nil).GetOracleDataBySpecID), arg0)
}

// GetSpecByID mocks base method
func (m *MockOracleService) GetSpecByID(arg0 string) (v1.OracleSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpecByID", arg0)
	ret0, _ := ret[0].(v1.OracleSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpecByID indicates an expected call of GetSpecByID
func (mr *MockOracleServiceMockRecorder) GetSpecByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpecByID", reflect.TypeOf((*MockOracleService)(nil).GetSpecByID), arg0)
}

// GetSpecs mocks base method
func (m *MockOracleService) GetSpecs() []v1.OracleSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpecs")
	ret0, _ := ret[0].([]v1.OracleSpec)
	return ret0
}

// GetSpecs indicates an expected call of GetSpecs
func (mr *MockOracleServiceMockRecorder) GetSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpecs", reflect.TypeOf((*MockOracleService)(nil).GetSpecs))
}