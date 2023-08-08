// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	scheduler "github.com/deepsquare-io/the-grid/supervisor/pkg/job/scheduler"
	mock "github.com/stretchr/testify/mock"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler struct {
	mock.Mock
}

type Scheduler_Expecter struct {
	mock *mock.Mock
}

func (_m *Scheduler) EXPECT() *Scheduler_Expecter {
	return &Scheduler_Expecter{mock: &_m.Mock}
}

// CancelJob provides a mock function with given fields: ctx, req
func (_m *Scheduler) CancelJob(ctx context.Context, req *scheduler.CancelRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.CancelRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scheduler_CancelJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelJob'
type Scheduler_CancelJob_Call struct {
	*mock.Call
}

// CancelJob is a helper method to define mock.On call
//   - ctx context.Context
//   - req *scheduler.CancelRequest
func (_e *Scheduler_Expecter) CancelJob(ctx interface{}, req interface{}) *Scheduler_CancelJob_Call {
	return &Scheduler_CancelJob_Call{Call: _e.mock.On("CancelJob", ctx, req)}
}

func (_c *Scheduler_CancelJob_Call) Run(run func(ctx context.Context, req *scheduler.CancelRequest)) *Scheduler_CancelJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*scheduler.CancelRequest))
	})
	return _c
}

func (_c *Scheduler_CancelJob_Call) Return(_a0 error) *Scheduler_CancelJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Scheduler_CancelJob_Call) RunAndReturn(run func(context.Context, *scheduler.CancelRequest) error) *Scheduler_CancelJob_Call {
	_c.Call.Return(run)
	return _c
}

// FindCPUsPerNode provides a mock function with given fields: ctx
func (_m *Scheduler) FindCPUsPerNode(ctx context.Context) ([]uint64, error) {
	ret := _m.Called(ctx)

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uint64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindCPUsPerNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindCPUsPerNode'
type Scheduler_FindCPUsPerNode_Call struct {
	*mock.Call
}

// FindCPUsPerNode is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindCPUsPerNode(ctx interface{}) *Scheduler_FindCPUsPerNode_Call {
	return &Scheduler_FindCPUsPerNode_Call{Call: _e.mock.On("FindCPUsPerNode", ctx)}
}

func (_c *Scheduler_FindCPUsPerNode_Call) Run(run func(ctx context.Context)) *Scheduler_FindCPUsPerNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindCPUsPerNode_Call) Return(_a0 []uint64, _a1 error) *Scheduler_FindCPUsPerNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindCPUsPerNode_Call) RunAndReturn(run func(context.Context) ([]uint64, error)) *Scheduler_FindCPUsPerNode_Call {
	_c.Call.Return(run)
	return _c
}

// FindGPUsPerNode provides a mock function with given fields: ctx
func (_m *Scheduler) FindGPUsPerNode(ctx context.Context) ([]uint64, error) {
	ret := _m.Called(ctx)

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uint64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindGPUsPerNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindGPUsPerNode'
type Scheduler_FindGPUsPerNode_Call struct {
	*mock.Call
}

// FindGPUsPerNode is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindGPUsPerNode(ctx interface{}) *Scheduler_FindGPUsPerNode_Call {
	return &Scheduler_FindGPUsPerNode_Call{Call: _e.mock.On("FindGPUsPerNode", ctx)}
}

func (_c *Scheduler_FindGPUsPerNode_Call) Run(run func(ctx context.Context)) *Scheduler_FindGPUsPerNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindGPUsPerNode_Call) Return(_a0 []uint64, _a1 error) *Scheduler_FindGPUsPerNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindGPUsPerNode_Call) RunAndReturn(run func(context.Context) ([]uint64, error)) *Scheduler_FindGPUsPerNode_Call {
	_c.Call.Return(run)
	return _c
}

// FindMemPerNode provides a mock function with given fields: ctx
func (_m *Scheduler) FindMemPerNode(ctx context.Context) ([]uint64, error) {
	ret := _m.Called(ctx)

	var r0 []uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []uint64); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindMemPerNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindMemPerNode'
type Scheduler_FindMemPerNode_Call struct {
	*mock.Call
}

// FindMemPerNode is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindMemPerNode(ctx interface{}) *Scheduler_FindMemPerNode_Call {
	return &Scheduler_FindMemPerNode_Call{Call: _e.mock.On("FindMemPerNode", ctx)}
}

func (_c *Scheduler_FindMemPerNode_Call) Run(run func(ctx context.Context)) *Scheduler_FindMemPerNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindMemPerNode_Call) Return(_a0 []uint64, _a1 error) *Scheduler_FindMemPerNode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindMemPerNode_Call) RunAndReturn(run func(context.Context) ([]uint64, error)) *Scheduler_FindMemPerNode_Call {
	_c.Call.Return(run)
	return _c
}

// FindRunningJobByName provides a mock function with given fields: ctx, req
func (_m *Scheduler) FindRunningJobByName(ctx context.Context, req *scheduler.FindRunningJobByNameRequest) (int, error) {
	ret := _m.Called(ctx, req)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.FindRunningJobByNameRequest) (int, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.FindRunningJobByNameRequest) int); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.FindRunningJobByNameRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindRunningJobByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindRunningJobByName'
type Scheduler_FindRunningJobByName_Call struct {
	*mock.Call
}

// FindRunningJobByName is a helper method to define mock.On call
//   - ctx context.Context
//   - req *scheduler.FindRunningJobByNameRequest
func (_e *Scheduler_Expecter) FindRunningJobByName(ctx interface{}, req interface{}) *Scheduler_FindRunningJobByName_Call {
	return &Scheduler_FindRunningJobByName_Call{Call: _e.mock.On("FindRunningJobByName", ctx, req)}
}

func (_c *Scheduler_FindRunningJobByName_Call) Run(run func(ctx context.Context, req *scheduler.FindRunningJobByNameRequest)) *Scheduler_FindRunningJobByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*scheduler.FindRunningJobByNameRequest))
	})
	return _c
}

func (_c *Scheduler_FindRunningJobByName_Call) Return(_a0 int, _a1 error) *Scheduler_FindRunningJobByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindRunningJobByName_Call) RunAndReturn(run func(context.Context, *scheduler.FindRunningJobByNameRequest) (int, error)) *Scheduler_FindRunningJobByName_Call {
	_c.Call.Return(run)
	return _c
}

// FindTotalCPUs provides a mock function with given fields: ctx
func (_m *Scheduler) FindTotalCPUs(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindTotalCPUs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTotalCPUs'
type Scheduler_FindTotalCPUs_Call struct {
	*mock.Call
}

// FindTotalCPUs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindTotalCPUs(ctx interface{}) *Scheduler_FindTotalCPUs_Call {
	return &Scheduler_FindTotalCPUs_Call{Call: _e.mock.On("FindTotalCPUs", ctx)}
}

func (_c *Scheduler_FindTotalCPUs_Call) Run(run func(ctx context.Context)) *Scheduler_FindTotalCPUs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindTotalCPUs_Call) Return(_a0 uint64, _a1 error) *Scheduler_FindTotalCPUs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindTotalCPUs_Call) RunAndReturn(run func(context.Context) (uint64, error)) *Scheduler_FindTotalCPUs_Call {
	_c.Call.Return(run)
	return _c
}

// FindTotalGPUs provides a mock function with given fields: ctx
func (_m *Scheduler) FindTotalGPUs(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindTotalGPUs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTotalGPUs'
type Scheduler_FindTotalGPUs_Call struct {
	*mock.Call
}

// FindTotalGPUs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindTotalGPUs(ctx interface{}) *Scheduler_FindTotalGPUs_Call {
	return &Scheduler_FindTotalGPUs_Call{Call: _e.mock.On("FindTotalGPUs", ctx)}
}

func (_c *Scheduler_FindTotalGPUs_Call) Run(run func(ctx context.Context)) *Scheduler_FindTotalGPUs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindTotalGPUs_Call) Return(_a0 uint64, _a1 error) *Scheduler_FindTotalGPUs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindTotalGPUs_Call) RunAndReturn(run func(context.Context) (uint64, error)) *Scheduler_FindTotalGPUs_Call {
	_c.Call.Return(run)
	return _c
}

// FindTotalMem provides a mock function with given fields: ctx
func (_m *Scheduler) FindTotalMem(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindTotalMem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTotalMem'
type Scheduler_FindTotalMem_Call struct {
	*mock.Call
}

// FindTotalMem is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindTotalMem(ctx interface{}) *Scheduler_FindTotalMem_Call {
	return &Scheduler_FindTotalMem_Call{Call: _e.mock.On("FindTotalMem", ctx)}
}

func (_c *Scheduler_FindTotalMem_Call) Run(run func(ctx context.Context)) *Scheduler_FindTotalMem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindTotalMem_Call) Return(_a0 uint64, _a1 error) *Scheduler_FindTotalMem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindTotalMem_Call) RunAndReturn(run func(context.Context) (uint64, error)) *Scheduler_FindTotalMem_Call {
	_c.Call.Return(run)
	return _c
}

// FindTotalNodes provides a mock function with given fields: ctx
func (_m *Scheduler) FindTotalNodes(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_FindTotalNodes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTotalNodes'
type Scheduler_FindTotalNodes_Call struct {
	*mock.Call
}

// FindTotalNodes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) FindTotalNodes(ctx interface{}) *Scheduler_FindTotalNodes_Call {
	return &Scheduler_FindTotalNodes_Call{Call: _e.mock.On("FindTotalNodes", ctx)}
}

func (_c *Scheduler_FindTotalNodes_Call) Run(run func(ctx context.Context)) *Scheduler_FindTotalNodes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_FindTotalNodes_Call) Return(_a0 uint64, _a1 error) *Scheduler_FindTotalNodes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_FindTotalNodes_Call) RunAndReturn(run func(context.Context) (uint64, error)) *Scheduler_FindTotalNodes_Call {
	_c.Call.Return(run)
	return _c
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *Scheduler) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scheduler_HealthCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthCheck'
type Scheduler_HealthCheck_Call struct {
	*mock.Call
}

// HealthCheck is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter) HealthCheck(ctx interface{}) *Scheduler_HealthCheck_Call {
	return &Scheduler_HealthCheck_Call{Call: _e.mock.On("HealthCheck", ctx)}
}

func (_c *Scheduler_HealthCheck_Call) Run(run func(ctx context.Context)) *Scheduler_HealthCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_HealthCheck_Call) Return(_a0 error) *Scheduler_HealthCheck_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Scheduler_HealthCheck_Call) RunAndReturn(run func(context.Context) error) *Scheduler_HealthCheck_Call {
	_c.Call.Return(run)
	return _c
}

// Submit provides a mock function with given fields: ctx, req
func (_m *Scheduler) Submit(ctx context.Context, req *scheduler.SubmitRequest) (string, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.SubmitRequest) (string, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.SubmitRequest) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scheduler.SubmitRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_Submit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Submit'
type Scheduler_Submit_Call struct {
	*mock.Call
}

// Submit is a helper method to define mock.On call
//   - ctx context.Context
//   - req *scheduler.SubmitRequest
func (_e *Scheduler_Expecter) Submit(ctx interface{}, req interface{}) *Scheduler_Submit_Call {
	return &Scheduler_Submit_Call{Call: _e.mock.On("Submit", ctx, req)}
}

func (_c *Scheduler_Submit_Call) Run(run func(ctx context.Context, req *scheduler.SubmitRequest)) *Scheduler_Submit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*scheduler.SubmitRequest))
	})
	return _c
}

func (_c *Scheduler_Submit_Call) Return(_a0 string, _a1 error) *Scheduler_Submit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_Submit_Call) RunAndReturn(run func(context.Context, *scheduler.SubmitRequest) (string, error)) *Scheduler_Submit_Call {
	_c.Call.Return(run)
	return _c
}

// TopUp provides a mock function with given fields: ctx, req
func (_m *Scheduler) TopUp(ctx context.Context, req *scheduler.TopUpRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *scheduler.TopUpRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scheduler_TopUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TopUp'
type Scheduler_TopUp_Call struct {
	*mock.Call
}

// TopUp is a helper method to define mock.On call
//   - ctx context.Context
//   - req *scheduler.TopUpRequest
func (_e *Scheduler_Expecter) TopUp(ctx interface{}, req interface{}) *Scheduler_TopUp_Call {
	return &Scheduler_TopUp_Call{Call: _e.mock.On("TopUp", ctx, req)}
}

func (_c *Scheduler_TopUp_Call) Run(run func(ctx context.Context, req *scheduler.TopUpRequest)) *Scheduler_TopUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*scheduler.TopUpRequest))
	})
	return _c
}

func (_c *Scheduler_TopUp_Call) Return(_a0 error) *Scheduler_TopUp_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Scheduler_TopUp_Call) RunAndReturn(run func(context.Context, *scheduler.TopUpRequest) error) *Scheduler_TopUp_Call {
	_c.Call.Return(run)
	return _c
}

// NewScheduler creates a new instance of Scheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Scheduler {
	mock := &Scheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
