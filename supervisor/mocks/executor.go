// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

type Executor_Expecter struct {
	mock *mock.Mock
}

func (_m *Executor) EXPECT() *Executor_Expecter {
	return &Executor_Expecter{mock: &_m.Mock}
}

// ExecAs provides a mock function with given fields: ctx, user, cmd
func (_m *Executor) ExecAs(ctx context.Context, user string, cmd string) (string, error) {
	ret := _m.Called(ctx, user, cmd)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, user, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, user, cmd)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, user, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Executor_ExecAs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecAs'
type Executor_ExecAs_Call struct {
	*mock.Call
}

// ExecAs is a helper method to define mock.On call
//   - ctx context.Context
//   - user string
//   - cmd string
func (_e *Executor_Expecter) ExecAs(ctx interface{}, user interface{}, cmd interface{}) *Executor_ExecAs_Call {
	return &Executor_ExecAs_Call{Call: _e.mock.On("ExecAs", ctx, user, cmd)}
}

func (_c *Executor_ExecAs_Call) Run(run func(ctx context.Context, user string, cmd string)) *Executor_ExecAs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *Executor_ExecAs_Call) Return(_a0 string, _a1 error) *Executor_ExecAs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Executor_ExecAs_Call) RunAndReturn(run func(context.Context, string, string) (string, error)) *Executor_ExecAs_Call {
	_c.Call.Return(run)
	return _c
}

// NewExecutor creates a new instance of Executor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Executor {
	mock := &Executor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}