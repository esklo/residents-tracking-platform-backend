// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: analytics/analytics.proto

package analytics

import (
	context "context"
	empty "github.com/esklo/residents-tracking-platform-backend/gen/proto/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnalyticsServiceClient is the client API for AnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticsServiceClient interface {
	RequestsPerTheme(ctx context.Context, in *RequestsPerThemeRequest, opts ...grpc.CallOption) (*RequestsPerThemeResponse, error)
	Stats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatsResponse, error)
	RequestsPerThemePerDate(ctx context.Context, in *RequestsPerThemeRequest, opts ...grpc.CallOption) (*RequestsPerThemePerDateResponse, error)
}

type analyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticsServiceClient(cc grpc.ClientConnInterface) AnalyticsServiceClient {
	return &analyticsServiceClient{cc}
}

func (c *analyticsServiceClient) RequestsPerTheme(ctx context.Context, in *RequestsPerThemeRequest, opts ...grpc.CallOption) (*RequestsPerThemeResponse, error) {
	out := new(RequestsPerThemeResponse)
	err := c.cc.Invoke(ctx, "/analytics.AnalyticsService/RequestsPerTheme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) Stats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/analytics.AnalyticsService/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyticsServiceClient) RequestsPerThemePerDate(ctx context.Context, in *RequestsPerThemeRequest, opts ...grpc.CallOption) (*RequestsPerThemePerDateResponse, error) {
	out := new(RequestsPerThemePerDateResponse)
	err := c.cc.Invoke(ctx, "/analytics.AnalyticsService/RequestsPerThemePerDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalyticsServiceServer is the server API for AnalyticsService service.
// All implementations must embed UnimplementedAnalyticsServiceServer
// for forward compatibility
type AnalyticsServiceServer interface {
	RequestsPerTheme(context.Context, *RequestsPerThemeRequest) (*RequestsPerThemeResponse, error)
	Stats(context.Context, *empty.Empty) (*StatsResponse, error)
	RequestsPerThemePerDate(context.Context, *RequestsPerThemeRequest) (*RequestsPerThemePerDateResponse, error)
	mustEmbedUnimplementedAnalyticsServiceServer()
}

// UnimplementedAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticsServiceServer struct {
}

func (UnimplementedAnalyticsServiceServer) RequestsPerTheme(context.Context, *RequestsPerThemeRequest) (*RequestsPerThemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsPerTheme not implemented")
}
func (UnimplementedAnalyticsServiceServer) Stats(context.Context, *empty.Empty) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedAnalyticsServiceServer) RequestsPerThemePerDate(context.Context, *RequestsPerThemeRequest) (*RequestsPerThemePerDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestsPerThemePerDate not implemented")
}
func (UnimplementedAnalyticsServiceServer) mustEmbedUnimplementedAnalyticsServiceServer() {}

// UnsafeAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticsServiceServer will
// result in compilation errors.
type UnsafeAnalyticsServiceServer interface {
	mustEmbedUnimplementedAnalyticsServiceServer()
}

func RegisterAnalyticsServiceServer(s grpc.ServiceRegistrar, srv AnalyticsServiceServer) {
	s.RegisterService(&AnalyticsService_ServiceDesc, srv)
}

func _AnalyticsService_RequestsPerTheme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsPerThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).RequestsPerTheme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.AnalyticsService/RequestsPerTheme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).RequestsPerTheme(ctx, req.(*RequestsPerThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.AnalyticsService/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).Stats(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnalyticsService_RequestsPerThemePerDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestsPerThemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyticsServiceServer).RequestsPerThemePerDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/analytics.AnalyticsService/RequestsPerThemePerDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyticsServiceServer).RequestsPerThemePerDate(ctx, req.(*RequestsPerThemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnalyticsService_ServiceDesc is the grpc.ServiceDesc for AnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.AnalyticsService",
	HandlerType: (*AnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestsPerTheme",
			Handler:    _AnalyticsService_RequestsPerTheme_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _AnalyticsService_Stats_Handler,
		},
		{
			MethodName: "RequestsPerThemePerDate",
			Handler:    _AnalyticsService_RequestsPerThemePerDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analytics/analytics.proto",
}
