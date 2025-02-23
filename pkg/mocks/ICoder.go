// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ICoder is an autogenerated mock type for the ICoder type
type ICoder[T interface{}] struct {
	mock.Mock
}

// Decode provides a mock function with given fields: data
func (_m *ICoder[T]) Decode(data []byte) (T, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (T, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) T); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encode provides a mock function with given fields: data
func (_m *ICoder[T]) Encode(data T) ([]byte, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(T) ([]byte, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(T) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(T) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICoder creates a new instance of ICoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICoder[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *ICoder[T] {
	mock := &ICoder[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
