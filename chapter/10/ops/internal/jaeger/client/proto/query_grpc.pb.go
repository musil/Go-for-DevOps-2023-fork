// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (QueryService_GetTraceClient, error)
	ArchiveTrace(ctx context.Context, in *ArchiveTraceRequest, opts ...grpc.CallOption) (*ArchiveTraceResponse, error)
	FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (QueryService_FindTracesClient, error)
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error)
	GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*GetDependenciesResponse, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (QueryService_GetTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryService_ServiceDesc.Streams[0], "/jaeger.api_v2.QueryService/GetTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceGetTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_GetTraceClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type queryServiceGetTraceClient struct {
	grpc.ClientStream
}

func (x *queryServiceGetTraceClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) ArchiveTrace(ctx context.Context, in *ArchiveTraceRequest, opts ...grpc.CallOption) (*ArchiveTraceResponse, error) {
	out := new(ArchiveTraceResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.QueryService/ArchiveTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (QueryService_FindTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryService_ServiceDesc.Streams[1], "/jaeger.api_v2.QueryService/FindTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceFindTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryService_FindTracesClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type queryServiceFindTracesClient struct {
	grpc.ClientStream
}

func (x *queryServiceFindTracesClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queryServiceClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.QueryService/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error) {
	out := new(GetOperationsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.QueryService/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*GetDependenciesResponse, error) {
	out := new(GetDependenciesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.api_v2.QueryService/GetDependencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations must embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	GetTrace(*GetTraceRequest, QueryService_GetTraceServer) error
	ArchiveTrace(context.Context, *ArchiveTraceRequest) (*ArchiveTraceResponse, error)
	FindTraces(*FindTracesRequest, QueryService_FindTracesServer) error
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error)
	GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error)
	mustEmbedUnimplementedQueryServiceServer()
}

// UnimplementedQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) GetTrace(*GetTraceRequest, QueryService_GetTraceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTrace not implemented")
}
func (UnimplementedQueryServiceServer) ArchiveTrace(context.Context, *ArchiveTraceRequest) (*ArchiveTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveTrace not implemented")
}
func (UnimplementedQueryServiceServer) FindTraces(*FindTracesRequest, QueryService_FindTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method FindTraces not implemented")
}
func (UnimplementedQueryServiceServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedQueryServiceServer) GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedQueryServiceServer) GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDependencies not implemented")
}
func (UnimplementedQueryServiceServer) mustEmbedUnimplementedQueryServiceServer() {}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_GetTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).GetTrace(m, &queryServiceGetTraceServer{stream})
}

type QueryService_GetTraceServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type queryServiceGetTraceServer struct {
	grpc.ServerStream
}

func (x *queryServiceGetTraceServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_ArchiveTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).ArchiveTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.QueryService/ArchiveTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).ArchiveTrace(ctx, req.(*ArchiveTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_FindTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindTracesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceServer).FindTraces(m, &queryServiceFindTracesServer{stream})
}

type QueryService_FindTracesServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type queryServiceFindTracesServer struct {
	grpc.ServerStream
}

func (x *queryServiceFindTracesServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _QueryService_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.QueryService/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.QueryService/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDependenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.api_v2.QueryService/GetDependencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetDependencies(ctx, req.(*GetDependenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.api_v2.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ArchiveTrace",
			Handler:    _QueryService_ArchiveTrace_Handler,
		},
		{
			MethodName: "GetServices",
			Handler:    _QueryService_GetServices_Handler,
		},
		{
			MethodName: "GetOperations",
			Handler:    _QueryService_GetOperations_Handler,
		},
		{
			MethodName: "GetDependencies",
			Handler:    _QueryService_GetDependencies_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTrace",
			Handler:       _QueryService_GetTrace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindTraces",
			Handler:       _QueryService_FindTraces_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query.proto",
}
