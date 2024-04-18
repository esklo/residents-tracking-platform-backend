// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: theme/theme.proto

package theme

import (
	context "context"
	empty "github.com/esklo/residents-tracking-platform/gen/proto/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ThemeService_Create_FullMethodName  = "/theme.ThemeService/Create"
	ThemeService_GetById_FullMethodName = "/theme.ThemeService/GetById"
	ThemeService_Get_FullMethodName     = "/theme.ThemeService/Get"
	ThemeService_Update_FullMethodName  = "/theme.ThemeService/Update"
	ThemeService_Delete_FullMethodName  = "/theme.ThemeService/Delete"
)

// ThemeServiceClient is the client API for ThemeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThemeServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Theme, error)
	GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*Theme, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *Theme, opts ...grpc.CallOption) (*Theme, error)
	Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type themeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThemeServiceClient(cc grpc.ClientConnInterface) ThemeServiceClient {
	return &themeServiceClient{cc}
}

func (c *themeServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Theme, error) {
	out := new(Theme)
	err := c.cc.Invoke(ctx, ThemeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) GetById(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*Theme, error) {
	out := new(Theme)
	err := c.cc.Invoke(ctx, ThemeService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ThemeService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) Update(ctx context.Context, in *Theme, opts ...grpc.CallOption) (*Theme, error) {
	out := new(Theme)
	err := c.cc.Invoke(ctx, ThemeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *themeServiceClient) Delete(ctx context.Context, in *ByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ThemeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThemeServiceServer is the server API for ThemeService service.
// All implementations must embed UnimplementedThemeServiceServer
// for forward compatibility
type ThemeServiceServer interface {
	Create(context.Context, *CreateRequest) (*Theme, error)
	GetById(context.Context, *ByIdRequest) (*Theme, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *Theme) (*Theme, error)
	Delete(context.Context, *ByIdRequest) (*empty.Empty, error)
	mustEmbedUnimplementedThemeServiceServer()
}

// UnimplementedThemeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThemeServiceServer struct {
}

func (UnimplementedThemeServiceServer) Create(context.Context, *CreateRequest) (*Theme, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedThemeServiceServer) GetById(context.Context, *ByIdRequest) (*Theme, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedThemeServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedThemeServiceServer) Update(context.Context, *Theme) (*Theme, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedThemeServiceServer) Delete(context.Context, *ByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedThemeServiceServer) mustEmbedUnimplementedThemeServiceServer() {}

// UnsafeThemeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThemeServiceServer will
// result in compilation errors.
type UnsafeThemeServiceServer interface {
	mustEmbedUnimplementedThemeServiceServer()
}

func RegisterThemeServiceServer(s grpc.ServiceRegistrar, srv ThemeServiceServer) {
	s.RegisterService(&ThemeService_ServiceDesc, srv)
}

func _ThemeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).GetById(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Theme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).Update(ctx, req.(*Theme))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThemeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThemeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThemeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThemeServiceServer).Delete(ctx, req.(*ByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThemeService_ServiceDesc is the grpc.ServiceDesc for ThemeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThemeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "theme.ThemeService",
	HandlerType: (*ThemeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ThemeService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ThemeService_GetById_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ThemeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ThemeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ThemeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "theme/theme.proto",
}
