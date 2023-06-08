// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: logger/v1alpha1/log.proto

package loggerv1alpha1

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

// LoggerAPIClient is the client API for LoggerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerAPIClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (LoggerAPI_WriteClient, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (LoggerAPI_ReadClient, error)
	WatchList(ctx context.Context, in *WatchListRequest, opts ...grpc.CallOption) (LoggerAPI_WatchListClient, error)
}

type loggerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerAPIClient(cc grpc.ClientConnInterface) LoggerAPIClient {
	return &loggerAPIClient{cc}
}

func (c *loggerAPIClient) Write(ctx context.Context, opts ...grpc.CallOption) (LoggerAPI_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggerAPI_ServiceDesc.Streams[0], "/logger.v1alpha1.LoggerAPI/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggerAPIWriteClient{stream}
	return x, nil
}

type LoggerAPI_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type loggerAPIWriteClient struct {
	grpc.ClientStream
}

func (x *loggerAPIWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loggerAPIWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loggerAPIClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (LoggerAPI_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggerAPI_ServiceDesc.Streams[1], "/logger.v1alpha1.LoggerAPI/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggerAPIReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoggerAPI_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type loggerAPIReadClient struct {
	grpc.ClientStream
}

func (x *loggerAPIReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *loggerAPIClient) WatchList(ctx context.Context, in *WatchListRequest, opts ...grpc.CallOption) (LoggerAPI_WatchListClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggerAPI_ServiceDesc.Streams[2], "/logger.v1alpha1.LoggerAPI/WatchList", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggerAPIWatchListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoggerAPI_WatchListClient interface {
	Recv() (*WatchListResponse, error)
	grpc.ClientStream
}

type loggerAPIWatchListClient struct {
	grpc.ClientStream
}

func (x *loggerAPIWatchListClient) Recv() (*WatchListResponse, error) {
	m := new(WatchListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggerAPIServer is the server API for LoggerAPI service.
// All implementations must embed UnimplementedLoggerAPIServer
// for forward compatibility
type LoggerAPIServer interface {
	Write(LoggerAPI_WriteServer) error
	Read(*ReadRequest, LoggerAPI_ReadServer) error
	WatchList(*WatchListRequest, LoggerAPI_WatchListServer) error
	mustEmbedUnimplementedLoggerAPIServer()
}

// UnimplementedLoggerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerAPIServer struct {
}

func (UnimplementedLoggerAPIServer) Write(LoggerAPI_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedLoggerAPIServer) Read(*ReadRequest, LoggerAPI_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedLoggerAPIServer) WatchList(*WatchListRequest, LoggerAPI_WatchListServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchList not implemented")
}
func (UnimplementedLoggerAPIServer) mustEmbedUnimplementedLoggerAPIServer() {}

// UnsafeLoggerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerAPIServer will
// result in compilation errors.
type UnsafeLoggerAPIServer interface {
	mustEmbedUnimplementedLoggerAPIServer()
}

func RegisterLoggerAPIServer(s grpc.ServiceRegistrar, srv LoggerAPIServer) {
	s.RegisterService(&LoggerAPI_ServiceDesc, srv)
}

func _LoggerAPI_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoggerAPIServer).Write(&loggerAPIWriteServer{stream})
}

type LoggerAPI_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type loggerAPIWriteServer struct {
	grpc.ServerStream
}

func (x *loggerAPIWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loggerAPIWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LoggerAPI_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoggerAPIServer).Read(m, &loggerAPIReadServer{stream})
}

type LoggerAPI_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type loggerAPIReadServer struct {
	grpc.ServerStream
}

func (x *loggerAPIReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LoggerAPI_WatchList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoggerAPIServer).WatchList(m, &loggerAPIWatchListServer{stream})
}

type LoggerAPI_WatchListServer interface {
	Send(*WatchListResponse) error
	grpc.ServerStream
}

type loggerAPIWatchListServer struct {
	grpc.ServerStream
}

func (x *loggerAPIWatchListServer) Send(m *WatchListResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LoggerAPI_ServiceDesc is the grpc.ServiceDesc for LoggerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.v1alpha1.LoggerAPI",
	HandlerType: (*LoggerAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _LoggerAPI_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _LoggerAPI_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchList",
			Handler:       _LoggerAPI_WatchList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logger/v1alpha1/log.proto",
}