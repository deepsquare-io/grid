// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeSBatchAPIServer is an autogenerated mock type for the UnsafeSBatchAPIServer type
type UnsafeSBatchAPIServer struct {
	mock.Mock
}

// mustEmbedUnimplementedSBatchAPIServer provides a mock function with given fields:
func (_m *UnsafeSBatchAPIServer) mustEmbedUnimplementedSBatchAPIServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeSBatchAPIServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeSBatchAPIServer creates a new instance of UnsafeSBatchAPIServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeSBatchAPIServer(t mockConstructorTestingTNewUnsafeSBatchAPIServer) *UnsafeSBatchAPIServer {
	mock := &UnsafeSBatchAPIServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}