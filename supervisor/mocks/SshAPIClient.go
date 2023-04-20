// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
)

// SshAPIClient is an autogenerated mock type for the SshAPIClient type
type SshAPIClient struct {
	mock.Mock
}

// FetchAuthorizedKeys provides a mock function with given fields: ctx, in, opts
func (_m *SshAPIClient) FetchAuthorizedKeys(ctx context.Context, in *supervisorv1alpha1.FetchAuthorizedKeysRequest, opts ...grpc.CallOption) (*supervisorv1alpha1.FetchAuthorizedKeysResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *supervisorv1alpha1.FetchAuthorizedKeysResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *supervisorv1alpha1.FetchAuthorizedKeysRequest, ...grpc.CallOption) (*supervisorv1alpha1.FetchAuthorizedKeysResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *supervisorv1alpha1.FetchAuthorizedKeysRequest, ...grpc.CallOption) *supervisorv1alpha1.FetchAuthorizedKeysResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*supervisorv1alpha1.FetchAuthorizedKeysResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *supervisorv1alpha1.FetchAuthorizedKeysRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSshAPIClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewSshAPIClient creates a new instance of SshAPIClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSshAPIClient(t mockConstructorTestingTNewSshAPIClient) *SshAPIClient {
	mock := &SshAPIClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
