// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetCandles provides a mock function with given fields: ctx, market, sinceTimestamp, interval
func (_m *Service) GetCandles(ctx context.Context, market string, sinceTimestamp uint64, interval msg.Interval) ([]*msg.Candle, error) {
	ret := _m.Called(ctx, market, sinceTimestamp, interval)

	var r0 []*msg.Candle
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, msg.Interval) []*msg.Candle); ok {
		r0 = rf(ctx, market, sinceTimestamp, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*msg.Candle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, msg.Interval) error); ok {
		r1 = rf(ctx, market, sinceTimestamp, interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ObserveCandles provides a mock function with given fields: ctx, market, interval
func (_m *Service) ObserveCandles(ctx context.Context, market *string, interval *msg.Interval) (<-chan msg.Candle, uint64) {
	ret := _m.Called(ctx, market, interval)

	var r0 <-chan msg.Candle
	if rf, ok := ret.Get(0).(func(context.Context, *string, *msg.Interval) <-chan msg.Candle); ok {
		r0 = rf(ctx, market, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan msg.Candle)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(context.Context, *string, *msg.Interval) uint64); ok {
		r1 = rf(ctx, market, interval)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	return r0, r1
}