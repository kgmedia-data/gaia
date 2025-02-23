// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	msg "github.com/kgmedia-data/gaia/pkg/msg"
	mock "github.com/stretchr/testify/mock"
)

// IPublisher is an autogenerated mock type for the IPublisher type
type IPublisher[T interface{}] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *IPublisher[T]) Close() error {
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

// Publish provides a mock function with given fields: message
func (_m *IPublisher[T]) Publish(message msg.Message[T]) error {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(msg.Message[T]) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIPublisher creates a new instance of IPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPublisher[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *IPublisher[T] {
	mock := &IPublisher[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
