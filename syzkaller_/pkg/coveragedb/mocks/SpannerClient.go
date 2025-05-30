// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	spanner "cloud.google.com/go/spanner"
	mock "github.com/stretchr/testify/mock"

	spannerclient "github.com/google/syzkaller/pkg/coveragedb/spannerclient"

	time "time"
)

// SpannerClient is an autogenerated mock type for the SpannerClient type
type SpannerClient struct {
	mock.Mock
}

// Apply provides a mock function with given fields: ctx, ms, opts
func (_m *SpannerClient) Apply(ctx context.Context, ms []*spanner.Mutation, opts ...spanner.ApplyOption) (time.Time, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ms)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Apply")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*spanner.Mutation, ...spanner.ApplyOption) (time.Time, error)); ok {
		return rf(ctx, ms, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*spanner.Mutation, ...spanner.ApplyOption) time.Time); ok {
		r0 = rf(ctx, ms, opts...)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*spanner.Mutation, ...spanner.ApplyOption) error); ok {
		r1 = rf(ctx, ms, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with no fields
func (_m *SpannerClient) Close() {
	_m.Called()
}

// Single provides a mock function with no fields
func (_m *SpannerClient) Single() spannerclient.ReadOnlyTransaction {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Single")
	}

	var r0 spannerclient.ReadOnlyTransaction
	if rf, ok := ret.Get(0).(func() spannerclient.ReadOnlyTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spannerclient.ReadOnlyTransaction)
		}
	}

	return r0
}

// NewSpannerClient creates a new instance of SpannerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSpannerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SpannerClient {
	mock := &SpannerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
