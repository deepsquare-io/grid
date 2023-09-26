// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocksbatch

import (
	context "context"

	sbatch "github.com/deepsquare-io/grid/supervisor/pkg/sbatch"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Fetch provides a mock function with given fields: ctx, hash
func (_m *Client) Fetch(ctx context.Context, hash string) (sbatch.FetchResponse, error) {
	ret := _m.Called(ctx, hash)

	var r0 sbatch.FetchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sbatch.FetchResponse, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sbatch.FetchResponse); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(sbatch.FetchResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type Client_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
//   - ctx context.Context
//   - hash string
func (_e *Client_Expecter) Fetch(ctx interface{}, hash interface{}) *Client_Fetch_Call {
	return &Client_Fetch_Call{Call: _e.mock.On("Fetch", ctx, hash)}
}

func (_c *Client_Fetch_Call) Run(run func(ctx context.Context, hash string)) *Client_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Client_Fetch_Call) Return(_a0 sbatch.FetchResponse, _a1 error) *Client_Fetch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Fetch_Call) RunAndReturn(run func(context.Context, string) (sbatch.FetchResponse, error)) *Client_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *Client) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_HealthCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthCheck'
type Client_HealthCheck_Call struct {
	*mock.Call
}

// HealthCheck is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) HealthCheck(ctx interface{}) *Client_HealthCheck_Call {
	return &Client_HealthCheck_Call{Call: _e.mock.On("HealthCheck", ctx)}
}

func (_c *Client_HealthCheck_Call) Run(run func(ctx context.Context)) *Client_HealthCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_HealthCheck_Call) Return(_a0 error) *Client_HealthCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_HealthCheck_Call) RunAndReturn(run func(context.Context) error) *Client_HealthCheck_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
