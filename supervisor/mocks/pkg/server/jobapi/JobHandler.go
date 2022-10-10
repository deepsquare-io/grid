// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// JobHandler is an autogenerated mock type for the JobHandler type
type JobHandler struct {
	mock.Mock
}

// FailJob provides a mock function with given fields: ctx, jobID
func (_m *JobHandler) FailJob(ctx context.Context, jobID [32]byte) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FinishJob provides a mock function with given fields: ctx, jobID, jobDuration
func (_m *JobHandler) FinishJob(ctx context.Context, jobID [32]byte, jobDuration uint64) error {
	ret := _m.Called(ctx, jobID, jobDuration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte, uint64) error); ok {
		r0 = rf(ctx, jobID, jobDuration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartJob provides a mock function with given fields: ctx, jobID
func (_m *JobHandler) StartJob(ctx context.Context, jobID [32]byte) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewJobHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewJobHandler creates a new instance of JobHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJobHandler(t mockConstructorTestingTNewJobHandler) *JobHandler {
	mock := &JobHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
