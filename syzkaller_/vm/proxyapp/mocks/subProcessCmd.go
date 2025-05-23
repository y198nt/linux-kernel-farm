// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// subProcessCmd is an autogenerated mock type for the subProcessCmd type
type subProcessCmd struct {
	mock.Mock
}

// Start provides a mock function with no fields
func (_m *subProcessCmd) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StderrPipe provides a mock function with no fields
func (_m *subProcessCmd) StderrPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StderrPipe")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StdinPipe provides a mock function with no fields
func (_m *subProcessCmd) StdinPipe() (io.WriteCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StdinPipe")
	}

	var r0 io.WriteCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.WriteCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.WriteCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.WriteCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StdoutPipe provides a mock function with no fields
func (_m *subProcessCmd) StdoutPipe() (io.ReadCloser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for StdoutPipe")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func() (io.ReadCloser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() io.ReadCloser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Wait provides a mock function with no fields
func (_m *subProcessCmd) Wait() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Wait")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newSubProcessCmd creates a new instance of subProcessCmd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newSubProcessCmd(t interface {
	mock.TestingT
	Cleanup(func())
}) *subProcessCmd {
	mock := &subProcessCmd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
