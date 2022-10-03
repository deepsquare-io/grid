// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	supervisorv1alpha1 "github.com/deepsquare-io/the-grid/supervisor/gen/go/supervisor/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// SshAPIServer is an autogenerated mock type for the SshAPIServer type
type SshAPIServer struct {
	mock.Mock
}

// FetchAuthorizedKeys provides a mock function with given fields: _a0, _a1
func (_m *SshAPIServer) FetchAuthorizedKeys(_a0 context.Context, _a1 *supervisorv1alpha1.FetchAuthorizedKeysRequest) (*supervisorv1alpha1.FetchAuthorizedKeysResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *supervisorv1alpha1.FetchAuthorizedKeysResponse
	if rf, ok := ret.Get(0).(func(context.Context, *supervisorv1alpha1.FetchAuthorizedKeysRequest) *supervisorv1alpha1.FetchAuthorizedKeysResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*supervisorv1alpha1.FetchAuthorizedKeysResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *supervisorv1alpha1.FetchAuthorizedKeysRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedSshAPIServer provides a mock function with given fields:
func (_m *SshAPIServer) mustEmbedUnimplementedSshAPIServer() {
	_m.Called()
}

type mockConstructorTestingTNewSshAPIServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewSshAPIServer creates a new instance of SshAPIServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSshAPIServer(t mockConstructorTestingTNewSshAPIServer) *SshAPIServer {
	mock := &SshAPIServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}