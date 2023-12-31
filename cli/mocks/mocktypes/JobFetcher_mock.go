// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocktypes

import (
	context "context"

	types "github.com/deepsquare-io/grid/cli/types"
	mock "github.com/stretchr/testify/mock"
)

// JobFetcher is an autogenerated mock type for the JobFetcher type
type JobFetcher struct {
	mock.Mock
}

type JobFetcher_Expecter struct {
	mock *mock.Mock
}

func (_m *JobFetcher) EXPECT() *JobFetcher_Expecter {
	return &JobFetcher_Expecter{mock: &_m.Mock}
}

// GetJob provides a mock function with given fields: ctx, id
func (_m *JobFetcher) GetJob(ctx context.Context, id [32]byte) (types.Job, error) {
	ret := _m.Called(ctx, id)

	var r0 types.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) (types.Job, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) types.Job); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [32]byte) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JobFetcher_GetJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJob'
type JobFetcher_GetJob_Call struct {
	*mock.Call
}

// GetJob is a helper method to define mock.On call
//   - ctx context.Context
//   - id [32]byte
func (_e *JobFetcher_Expecter) GetJob(ctx interface{}, id interface{}) *JobFetcher_GetJob_Call {
	return &JobFetcher_GetJob_Call{Call: _e.mock.On("GetJob", ctx, id)}
}

func (_c *JobFetcher_GetJob_Call) Run(run func(ctx context.Context, id [32]byte)) *JobFetcher_GetJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([32]byte))
	})
	return _c
}

func (_c *JobFetcher_GetJob_Call) Return(_a0 types.Job, _a1 error) *JobFetcher_GetJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JobFetcher_GetJob_Call) RunAndReturn(run func(context.Context, [32]byte) (types.Job, error)) *JobFetcher_GetJob_Call {
	_c.Call.Return(run)
	return _c
}

// GetJobs provides a mock function with given fields: ctx
func (_m *JobFetcher) GetJobs(ctx context.Context) (types.JobLazyIterator, error) {
	ret := _m.Called(ctx)

	var r0 types.JobLazyIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (types.JobLazyIterator, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) types.JobLazyIterator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.JobLazyIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JobFetcher_GetJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJobs'
type JobFetcher_GetJobs_Call struct {
	*mock.Call
}

// GetJobs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *JobFetcher_Expecter) GetJobs(ctx interface{}) *JobFetcher_GetJobs_Call {
	return &JobFetcher_GetJobs_Call{Call: _e.mock.On("GetJobs", ctx)}
}

func (_c *JobFetcher_GetJobs_Call) Run(run func(ctx context.Context)) *JobFetcher_GetJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *JobFetcher_GetJobs_Call) Return(_a0 types.JobLazyIterator, _a1 error) *JobFetcher_GetJobs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *JobFetcher_GetJobs_Call) RunAndReturn(run func(context.Context) (types.JobLazyIterator, error)) *JobFetcher_GetJobs_Call {
	_c.Call.Return(run)
	return _c
}

// NewJobFetcher creates a new instance of JobFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJobFetcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *JobFetcher {
	mock := &JobFetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
