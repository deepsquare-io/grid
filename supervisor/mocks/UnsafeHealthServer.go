// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeHealthServer is an autogenerated mock type for the UnsafeHealthServer type
type UnsafeHealthServer struct {
	mock.Mock
}

// mustEmbedUnimplementedHealthServer provides a mock function with given fields:
func (_m *UnsafeHealthServer) mustEmbedUnimplementedHealthServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeHealthServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeHealthServer creates a new instance of UnsafeHealthServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeHealthServer(t mockConstructorTestingTNewUnsafeHealthServer) *UnsafeHealthServer {
	mock := &UnsafeHealthServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
