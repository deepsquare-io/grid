// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: supervisor/v1alpha1/ssh.proto

package supervisorv1alpha1

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

// SshAPIClient is the client API for SshAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SshAPIClient interface {
	FetchAuthorizedKeys(ctx context.Context, in *FetchAuthorizedKeysRequest, opts ...grpc.CallOption) (*FetchAuthorizedKeysResponse, error)
}

type sshAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSshAPIClient(cc grpc.ClientConnInterface) SshAPIClient {
	return &sshAPIClient{cc}
}

func (c *sshAPIClient) FetchAuthorizedKeys(ctx context.Context, in *FetchAuthorizedKeysRequest, opts ...grpc.CallOption) (*FetchAuthorizedKeysResponse, error) {
	out := new(FetchAuthorizedKeysResponse)
	err := c.cc.Invoke(ctx, "/supervisor.v1alpha1.SshAPI/FetchAuthorizedKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SshAPIServer is the server API for SshAPI service.
// All implementations must embed UnimplementedSshAPIServer
// for forward compatibility
type SshAPIServer interface {
	FetchAuthorizedKeys(context.Context, *FetchAuthorizedKeysRequest) (*FetchAuthorizedKeysResponse, error)
	mustEmbedUnimplementedSshAPIServer()
}

// UnimplementedSshAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSshAPIServer struct {
}

func (UnimplementedSshAPIServer) FetchAuthorizedKeys(context.Context, *FetchAuthorizedKeysRequest) (*FetchAuthorizedKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAuthorizedKeys not implemented")
}
func (UnimplementedSshAPIServer) mustEmbedUnimplementedSshAPIServer() {}

// UnsafeSshAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SshAPIServer will
// result in compilation errors.
type UnsafeSshAPIServer interface {
	mustEmbedUnimplementedSshAPIServer()
}

func RegisterSshAPIServer(s grpc.ServiceRegistrar, srv SshAPIServer) {
	s.RegisterService(&SshAPI_ServiceDesc, srv)
}

func _SshAPI_FetchAuthorizedKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAuthorizedKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SshAPIServer).FetchAuthorizedKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1alpha1.SshAPI/FetchAuthorizedKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SshAPIServer).FetchAuthorizedKeys(ctx, req.(*FetchAuthorizedKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SshAPI_ServiceDesc is the grpc.ServiceDesc for SshAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SshAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.v1alpha1.SshAPI",
	HandlerType: (*SshAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAuthorizedKeys",
			Handler:    _SshAPI_FetchAuthorizedKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor/v1alpha1/ssh.proto",
}
