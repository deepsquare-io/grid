// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	metascheduler "github.com/deepsquare-io/the-grid/supervisor/pkg/metascheduler"

	metaschedulerabi "github.com/deepsquare-io/the-grid/supervisor/generated/abi/metascheduler"

	mock "github.com/stretchr/testify/mock"
)

// MetaScheduler is an autogenerated mock type for the MetaScheduler type
type MetaScheduler struct {
	mock.Mock
}

type MetaScheduler_Expecter struct {
	mock *mock.Mock
}

func (_m *MetaScheduler) EXPECT() *MetaScheduler_Expecter {
	return &MetaScheduler_Expecter{mock: &_m.Mock}
}

// Claim provides a mock function with given fields: ctx
func (_m *MetaScheduler) Claim(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaScheduler_Claim_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Claim'
type MetaScheduler_Claim_Call struct {
	*mock.Call
}

// Claim is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetaScheduler_Expecter) Claim(ctx interface{}) *MetaScheduler_Claim_Call {
	return &MetaScheduler_Claim_Call{Call: _e.mock.On("Claim", ctx)}
}

func (_c *MetaScheduler_Claim_Call) Run(run func(ctx context.Context)) *MetaScheduler_Claim_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetaScheduler_Claim_Call) Return(_a0 error) *MetaScheduler_Claim_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_Claim_Call) RunAndReturn(run func(context.Context) error) *MetaScheduler_Claim_Call {
	_c.Call.Return(run)
	return _c
}

// ClaimCancelling provides a mock function with given fields: ctx
func (_m *MetaScheduler) ClaimCancelling(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaScheduler_ClaimCancelling_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClaimCancelling'
type MetaScheduler_ClaimCancelling_Call struct {
	*mock.Call
}

// ClaimCancelling is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetaScheduler_Expecter) ClaimCancelling(ctx interface{}) *MetaScheduler_ClaimCancelling_Call {
	return &MetaScheduler_ClaimCancelling_Call{Call: _e.mock.On("ClaimCancelling", ctx)}
}

func (_c *MetaScheduler_ClaimCancelling_Call) Run(run func(ctx context.Context)) *MetaScheduler_ClaimCancelling_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetaScheduler_ClaimCancelling_Call) Return(_a0 error) *MetaScheduler_ClaimCancelling_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_ClaimCancelling_Call) RunAndReturn(run func(context.Context) error) *MetaScheduler_ClaimCancelling_Call {
	_c.Call.Return(run)
	return _c
}

// ClaimTopUp provides a mock function with given fields: ctx
func (_m *MetaScheduler) ClaimTopUp(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaScheduler_ClaimTopUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClaimTopUp'
type MetaScheduler_ClaimTopUp_Call struct {
	*mock.Call
}

// ClaimTopUp is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MetaScheduler_Expecter) ClaimTopUp(ctx interface{}) *MetaScheduler_ClaimTopUp_Call {
	return &MetaScheduler_ClaimTopUp_Call{Call: _e.mock.On("ClaimTopUp", ctx)}
}

func (_c *MetaScheduler_ClaimTopUp_Call) Run(run func(ctx context.Context)) *MetaScheduler_ClaimTopUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MetaScheduler_ClaimTopUp_Call) Return(_a0 error) *MetaScheduler_ClaimTopUp_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_ClaimTopUp_Call) RunAndReturn(run func(context.Context) error) *MetaScheduler_ClaimTopUp_Call {
	_c.Call.Return(run)
	return _c
}

// GetJobStatus provides a mock function with given fields: ctx, jobID
func (_m *MetaScheduler) GetJobStatus(ctx context.Context, jobID [32]byte) (metascheduler.JobStatus, error) {
	ret := _m.Called(ctx, jobID)

	var r0 metascheduler.JobStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) (metascheduler.JobStatus, error)); ok {
		return rf(ctx, jobID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) metascheduler.JobStatus); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Get(0).(metascheduler.JobStatus)
	}

	if rf, ok := ret.Get(1).(func(context.Context, [32]byte) error); ok {
		r1 = rf(ctx, jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaScheduler_GetJobStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJobStatus'
type MetaScheduler_GetJobStatus_Call struct {
	*mock.Call
}

// GetJobStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - jobID [32]byte
func (_e *MetaScheduler_Expecter) GetJobStatus(ctx interface{}, jobID interface{}) *MetaScheduler_GetJobStatus_Call {
	return &MetaScheduler_GetJobStatus_Call{Call: _e.mock.On("GetJobStatus", ctx, jobID)}
}

func (_c *MetaScheduler_GetJobStatus_Call) Run(run func(ctx context.Context, jobID [32]byte)) *MetaScheduler_GetJobStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([32]byte))
	})
	return _c
}

func (_c *MetaScheduler_GetJobStatus_Call) Return(_a0 metascheduler.JobStatus, _a1 error) *MetaScheduler_GetJobStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaScheduler_GetJobStatus_Call) RunAndReturn(run func(context.Context, [32]byte) (metascheduler.JobStatus, error)) *MetaScheduler_GetJobStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetProviderAddress provides a mock function with given fields:
func (_m *MetaScheduler) GetProviderAddress() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// MetaScheduler_GetProviderAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProviderAddress'
type MetaScheduler_GetProviderAddress_Call struct {
	*mock.Call
}

// GetProviderAddress is a helper method to define mock.On call
func (_e *MetaScheduler_Expecter) GetProviderAddress() *MetaScheduler_GetProviderAddress_Call {
	return &MetaScheduler_GetProviderAddress_Call{Call: _e.mock.On("GetProviderAddress")}
}

func (_c *MetaScheduler_GetProviderAddress_Call) Run(run func()) *MetaScheduler_GetProviderAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MetaScheduler_GetProviderAddress_Call) Return(_a0 common.Address) *MetaScheduler_GetProviderAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_GetProviderAddress_Call) RunAndReturn(run func() common.Address) *MetaScheduler_GetProviderAddress_Call {
	_c.Call.Return(run)
	return _c
}

// RefuseJob provides a mock function with given fields: ctx, jobID
func (_m *MetaScheduler) RefuseJob(ctx context.Context, jobID [32]byte) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaScheduler_RefuseJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefuseJob'
type MetaScheduler_RefuseJob_Call struct {
	*mock.Call
}

// RefuseJob is a helper method to define mock.On call
//   - ctx context.Context
//   - jobID [32]byte
func (_e *MetaScheduler_Expecter) RefuseJob(ctx interface{}, jobID interface{}) *MetaScheduler_RefuseJob_Call {
	return &MetaScheduler_RefuseJob_Call{Call: _e.mock.On("RefuseJob", ctx, jobID)}
}

func (_c *MetaScheduler_RefuseJob_Call) Run(run func(ctx context.Context, jobID [32]byte)) *MetaScheduler_RefuseJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([32]byte))
	})
	return _c
}

func (_c *MetaScheduler_RefuseJob_Call) Return(_a0 error) *MetaScheduler_RefuseJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_RefuseJob_Call) RunAndReturn(run func(context.Context, [32]byte) error) *MetaScheduler_RefuseJob_Call {
	_c.Call.Return(run)
	return _c
}

// SetJobStatus provides a mock function with given fields: ctx, jobID, status, jobDurationMinute
func (_m *MetaScheduler) SetJobStatus(ctx context.Context, jobID [32]byte, status metascheduler.JobStatus, jobDurationMinute uint64) error {
	ret := _m.Called(ctx, jobID, status, jobDurationMinute)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte, metascheduler.JobStatus, uint64) error); ok {
		r0 = rf(ctx, jobID, status, jobDurationMinute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MetaScheduler_SetJobStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetJobStatus'
type MetaScheduler_SetJobStatus_Call struct {
	*mock.Call
}

// SetJobStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - jobID [32]byte
//   - status metascheduler.JobStatus
//   - jobDurationMinute uint64
func (_e *MetaScheduler_Expecter) SetJobStatus(ctx interface{}, jobID interface{}, status interface{}, jobDurationMinute interface{}) *MetaScheduler_SetJobStatus_Call {
	return &MetaScheduler_SetJobStatus_Call{Call: _e.mock.On("SetJobStatus", ctx, jobID, status, jobDurationMinute)}
}

func (_c *MetaScheduler_SetJobStatus_Call) Run(run func(ctx context.Context, jobID [32]byte, status metascheduler.JobStatus, jobDurationMinute uint64)) *MetaScheduler_SetJobStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([32]byte), args[2].(metascheduler.JobStatus), args[3].(uint64))
	})
	return _c
}

func (_c *MetaScheduler_SetJobStatus_Call) Return(_a0 error) *MetaScheduler_SetJobStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MetaScheduler_SetJobStatus_Call) RunAndReturn(run func(context.Context, [32]byte, metascheduler.JobStatus, uint64) error) *MetaScheduler_SetJobStatus_Call {
	_c.Call.Return(run)
	return _c
}

// WatchEvents provides a mock function with given fields: ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents
func (_m *MetaScheduler) WatchEvents(ctx context.Context, claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error) {
	ret := _m.Called(ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents)

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error)); ok {
		return rf(ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) event.Subscription); ok {
		r0 = rf(ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) error); ok {
		r1 = rf(ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaScheduler_WatchEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchEvents'
type MetaScheduler_WatchEvents_Call struct {
	*mock.Call
}

// WatchEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent
//   - claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent
//   - claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent
func (_e *MetaScheduler_Expecter) WatchEvents(ctx interface{}, claimNextTopUpJobEvents interface{}, claimNextCancellingJobEvents interface{}, claimJobEvents interface{}) *MetaScheduler_WatchEvents_Call {
	return &MetaScheduler_WatchEvents_Call{Call: _e.mock.On("WatchEvents", ctx, claimNextTopUpJobEvents, claimNextCancellingJobEvents, claimJobEvents)}
}

func (_c *MetaScheduler_WatchEvents_Call) Run(run func(ctx context.Context, claimNextTopUpJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, claimNextCancellingJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, claimJobEvents chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent)) *MetaScheduler_WatchEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent), args[2].(chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent), args[3].(chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent))
	})
	return _c
}

func (_c *MetaScheduler_WatchEvents_Call) Return(_a0 event.Subscription, _a1 error) *MetaScheduler_WatchEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetaScheduler_WatchEvents_Call) RunAndReturn(run func(context.Context, chan<- *metaschedulerabi.MetaSchedulerClaimNextTopUpJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimNextCancellingJobEvent, chan<- *metaschedulerabi.MetaSchedulerClaimJobEvent) (event.Subscription, error)) *MetaScheduler_WatchEvents_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetaScheduler creates a new instance of MetaScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetaScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetaScheduler {
	mock := &MetaScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
