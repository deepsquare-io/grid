// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: sbatchapi/v1alpha1/sbatch.proto

package sbatchapiv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthAPIClient interface {
	Challenge(ctx context.Context, in *ChallengeRequest, opts ...grpc.CallOption) (*ChallengeResponse, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) Challenge(ctx context.Context, in *ChallengeRequest, opts ...grpc.CallOption) (*ChallengeResponse, error) {
	out := new(ChallengeResponse)
	err := c.cc.Invoke(ctx, "/sbatchapi.v1alpha1.AuthAPI/Challenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
// All implementations must embed UnimplementedAuthAPIServer
// for forward compatibility
type AuthAPIServer interface {
	Challenge(context.Context, *ChallengeRequest) (*ChallengeResponse, error)
	mustEmbedUnimplementedAuthAPIServer()
}

// UnimplementedAuthAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (UnimplementedAuthAPIServer) Challenge(context.Context, *ChallengeRequest) (*ChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Challenge not implemented")
}
func (UnimplementedAuthAPIServer) mustEmbedUnimplementedAuthAPIServer() {}

// UnsafeAuthAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthAPIServer will
// result in compilation errors.
type UnsafeAuthAPIServer interface {
	mustEmbedUnimplementedAuthAPIServer()
}

func RegisterAuthAPIServer(s grpc.ServiceRegistrar, srv AuthAPIServer) {
	s.RegisterService(&AuthAPI_ServiceDesc, srv)
}

func _AuthAPI_Challenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).Challenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sbatchapi.v1alpha1.AuthAPI/Challenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).Challenge(ctx, req.(*ChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthAPI_ServiceDesc is the grpc.ServiceDesc for AuthAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sbatchapi.v1alpha1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Challenge",
			Handler:    _AuthAPI_Challenge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sbatchapi/v1alpha1/sbatch.proto",
}

// SBatchAPIClient is the client API for SBatchAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SBatchAPIClient interface {
	GetSBatch(ctx context.Context, in *GetSBatchRequest, opts ...grpc.CallOption) (*GetSBatchResponse, error)
}

type sBatchAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSBatchAPIClient(cc grpc.ClientConnInterface) SBatchAPIClient {
	return &sBatchAPIClient{cc}
}

func (c *sBatchAPIClient) GetSBatch(ctx context.Context, in *GetSBatchRequest, opts ...grpc.CallOption) (*GetSBatchResponse, error) {
	out := new(GetSBatchResponse)
	err := c.cc.Invoke(ctx, "/sbatchapi.v1alpha1.SBatchAPI/GetSBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SBatchAPIServer is the server API for SBatchAPI service.
// All implementations must embed UnimplementedSBatchAPIServer
// for forward compatibility
type SBatchAPIServer interface {
	GetSBatch(context.Context, *GetSBatchRequest) (*GetSBatchResponse, error)
	mustEmbedUnimplementedSBatchAPIServer()
}

// UnimplementedSBatchAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSBatchAPIServer struct {
}

func (UnimplementedSBatchAPIServer) GetSBatch(context.Context, *GetSBatchRequest) (*GetSBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSBatch not implemented")
}
func (UnimplementedSBatchAPIServer) mustEmbedUnimplementedSBatchAPIServer() {}

// UnsafeSBatchAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SBatchAPIServer will
// result in compilation errors.
type UnsafeSBatchAPIServer interface {
	mustEmbedUnimplementedSBatchAPIServer()
}

func RegisterSBatchAPIServer(s grpc.ServiceRegistrar, srv SBatchAPIServer) {
	s.RegisterService(&SBatchAPI_ServiceDesc, srv)
}

func _SBatchAPI_GetSBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SBatchAPIServer).GetSBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sbatchapi.v1alpha1.SBatchAPI/GetSBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SBatchAPIServer).GetSBatch(ctx, req.(*GetSBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SBatchAPI_ServiceDesc is the grpc.ServiceDesc for SBatchAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SBatchAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sbatchapi.v1alpha1.SBatchAPI",
	HandlerType: (*SBatchAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSBatch",
			Handler:    _SBatchAPI_GetSBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sbatchapi/v1alpha1/sbatch.proto",
}
