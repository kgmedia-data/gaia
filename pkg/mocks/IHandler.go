// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Start provides a mock function with given fields:
func (_m *IHandler) Start() error {
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

// Stop provides a mock function with given fields:
func (_m *IHandler) Stop() {
	_m.Called()
}

// NewIHandler creates a new instance of IHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHandler {
	mock := &IHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}