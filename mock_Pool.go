// Code generated by mockery v2.46.1. DO NOT EDIT.

package goredis

import (
	context "context"

	redigoredis "github.com/gomodule/redigo/redis"
	mock "github.com/stretchr/testify/mock"
)

// MockPool is an autogenerated mock type for the Pool type
type MockPool struct {
	mock.Mock
}

// Conn provides a mock function with given fields:
func (_m *MockPool) Conn() redigoredis.Conn {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Conn")
	}

	var r0 redigoredis.Conn
	if rf, ok := ret.Get(0).(func() redigoredis.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(redigoredis.Conn)
		}
	}

	return r0
}

// Do provides a mock function with given fields: command, args
func (_m *MockPool) Do(command string, args ...any) (any, error) {
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 any
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...any) (any, error)); ok {
		return rf(command, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...any) any); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Get(0).(any)
	}

	if rf, ok := ret.Get(1).(func(string, ...any) error); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoCtx provides a mock function with given fields: ctx, command, args
func (_m *MockPool) DoCtx(ctx context.Context, command string, args ...any) (any, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, command)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DoCtx")
	}

	var r0 any
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...any) (any, error)); ok {
		return rf(ctx, command, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...any) any); ok {
		r0 = rf(ctx, command, args...)
	} else {
		r0 = ret.Get(0).(any)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...any) error); ok {
		r1 = rf(ctx, command, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockPool creates a new instance of MockPool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPool(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPool {
	mock := &MockPool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}