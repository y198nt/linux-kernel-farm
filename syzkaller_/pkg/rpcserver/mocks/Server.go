// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dispatcher "github.com/google/syzkaller/vm/dispatcher"
	mock "github.com/stretchr/testify/mock"

	report "github.com/google/syzkaller/pkg/report"

	rpcserver "github.com/google/syzkaller/pkg/rpcserver"

	signal "github.com/google/syzkaller/pkg/signal"
)

// Server is an autogenerated mock type for the Server type
type Server struct {
	mock.Mock
}

// Close provides a mock function with no fields
func (_m *Server) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInstance provides a mock function with given fields: id, injectExec, updInfo
func (_m *Server) CreateInstance(id int, injectExec chan<- bool, updInfo dispatcher.UpdateInfo) chan error {
	ret := _m.Called(id, injectExec, updInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateInstance")
	}

	var r0 chan error
	if rf, ok := ret.Get(0).(func(int, chan<- bool, dispatcher.UpdateInfo) chan error); ok {
		r0 = rf(id, injectExec, updInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan error)
		}
	}

	return r0
}

// DistributeSignalDelta provides a mock function with given fields: plus
func (_m *Server) DistributeSignalDelta(plus signal.Signal) {
	_m.Called(plus)
}

// Listen provides a mock function with no fields
func (_m *Server) Listen() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Listen")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Port provides a mock function with no fields
func (_m *Server) Port() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Port")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Serve provides a mock function with given fields: _a0
func (_m *Server) Serve(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Serve")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShutdownInstance provides a mock function with given fields: id, crashed, extraExecs
func (_m *Server) ShutdownInstance(id int, crashed bool, extraExecs ...report.ExecutorInfo) ([]rpcserver.ExecRecord, []byte) {
	_va := make([]interface{}, len(extraExecs))
	for _i := range extraExecs {
		_va[_i] = extraExecs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id, crashed)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ShutdownInstance")
	}

	var r0 []rpcserver.ExecRecord
	var r1 []byte
	if rf, ok := ret.Get(0).(func(int, bool, ...report.ExecutorInfo) ([]rpcserver.ExecRecord, []byte)); ok {
		return rf(id, crashed, extraExecs...)
	}
	if rf, ok := ret.Get(0).(func(int, bool, ...report.ExecutorInfo) []rpcserver.ExecRecord); ok {
		r0 = rf(id, crashed, extraExecs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rpcserver.ExecRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(int, bool, ...report.ExecutorInfo) []byte); ok {
		r1 = rf(id, crashed, extraExecs...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	return r0, r1
}

// StopFuzzing provides a mock function with given fields: id
func (_m *Server) StopFuzzing(id int) {
	_m.Called(id)
}

// TriagedCorpus provides a mock function with no fields
func (_m *Server) TriagedCorpus() {
	_m.Called()
}

// NewServer creates a new instance of Server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Server {
	mock := &Server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
