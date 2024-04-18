// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: geo/geo.proto

package geo

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

const (
	GeoService_BuildingsByCoordinates_FullMethodName     = "/geo.GeoService/BuildingsByCoordinates"
	GeoService_GetAdministrativeDistricts_FullMethodName = "/geo.GeoService/GetAdministrativeDistricts"
	GeoService_GetDistricts_FullMethodName               = "/geo.GeoService/GetDistricts"
	GeoService_Suggest_FullMethodName                    = "/geo.GeoService/Suggest"
	GeoService_GeoLocate_FullMethodName                  = "/geo.GeoService/GeoLocate"
)

// GeoServiceClient is the client API for GeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoServiceClient interface {
	BuildingsByCoordinates(ctx context.Context, in *GeoPoint, opts ...grpc.CallOption) (*BuildingByCoordinatesResponse, error)
	GetAdministrativeDistricts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetDistrictsResponse, error)
	GetDistricts(ctx context.Context, in *GetDistrictsRequest, opts ...grpc.CallOption) (*GetDistrictsResponse, error)
	Suggest(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error)
	GeoLocate(ctx context.Context, in *GeoPoint, opts ...grpc.CallOption) (*GeoLocateResponse, error)
}

type geoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoServiceClient(cc grpc.ClientConnInterface) GeoServiceClient {
	return &geoServiceClient{cc}
}

func (c *geoServiceClient) BuildingsByCoordinates(ctx context.Context, in *GeoPoint, opts ...grpc.CallOption) (*BuildingByCoordinatesResponse, error) {
	out := new(BuildingByCoordinatesResponse)
	err := c.cc.Invoke(ctx, GeoService_BuildingsByCoordinates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) GetAdministrativeDistricts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetDistrictsResponse, error) {
	out := new(GetDistrictsResponse)
	err := c.cc.Invoke(ctx, GeoService_GetAdministrativeDistricts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) GetDistricts(ctx context.Context, in *GetDistrictsRequest, opts ...grpc.CallOption) (*GetDistrictsResponse, error) {
	out := new(GetDistrictsResponse)
	err := c.cc.Invoke(ctx, GeoService_GetDistricts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) Suggest(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error) {
	out := new(SuggestResponse)
	err := c.cc.Invoke(ctx, GeoService_Suggest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) GeoLocate(ctx context.Context, in *GeoPoint, opts ...grpc.CallOption) (*GeoLocateResponse, error) {
	out := new(GeoLocateResponse)
	err := c.cc.Invoke(ctx, GeoService_GeoLocate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServiceServer is the server API for GeoService service.
// All implementations must embed UnimplementedGeoServiceServer
// for forward compatibility
type GeoServiceServer interface {
	BuildingsByCoordinates(context.Context, *GeoPoint) (*BuildingByCoordinatesResponse, error)
	GetAdministrativeDistricts(context.Context, *empty.Empty) (*GetDistrictsResponse, error)
	GetDistricts(context.Context, *GetDistrictsRequest) (*GetDistrictsResponse, error)
	Suggest(context.Context, *SuggestRequest) (*SuggestResponse, error)
	GeoLocate(context.Context, *GeoPoint) (*GeoLocateResponse, error)
	mustEmbedUnimplementedGeoServiceServer()
}

// UnimplementedGeoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeoServiceServer struct {
}

func (UnimplementedGeoServiceServer) BuildingsByCoordinates(context.Context, *GeoPoint) (*BuildingByCoordinatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildingsByCoordinates not implemented")
}
func (UnimplementedGeoServiceServer) GetAdministrativeDistricts(context.Context, *empty.Empty) (*GetDistrictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdministrativeDistricts not implemented")
}
func (UnimplementedGeoServiceServer) GetDistricts(context.Context, *GetDistrictsRequest) (*GetDistrictsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistricts not implemented")
}
func (UnimplementedGeoServiceServer) Suggest(context.Context, *SuggestRequest) (*SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggest not implemented")
}
func (UnimplementedGeoServiceServer) GeoLocate(context.Context, *GeoPoint) (*GeoLocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocate not implemented")
}
func (UnimplementedGeoServiceServer) mustEmbedUnimplementedGeoServiceServer() {}

// UnsafeGeoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoServiceServer will
// result in compilation errors.
type UnsafeGeoServiceServer interface {
	mustEmbedUnimplementedGeoServiceServer()
}

func RegisterGeoServiceServer(s grpc.ServiceRegistrar, srv GeoServiceServer) {
	s.RegisterService(&GeoService_ServiceDesc, srv)
}

func _GeoService_BuildingsByCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoPoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).BuildingsByCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_BuildingsByCoordinates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).BuildingsByCoordinates(ctx, req.(*GeoPoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_GetAdministrativeDistricts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GetAdministrativeDistricts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_GetAdministrativeDistricts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GetAdministrativeDistricts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_GetDistricts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistrictsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GetDistricts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_GetDistricts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GetDistricts(ctx, req.(*GetDistrictsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_Suggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).Suggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_Suggest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).Suggest(ctx, req.(*SuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_GeoLocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoPoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GeoLocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoService_GeoLocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GeoLocate(ctx, req.(*GeoPoint))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoService_ServiceDesc is the grpc.ServiceDesc for GeoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "geo.GeoService",
	HandlerType: (*GeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildingsByCoordinates",
			Handler:    _GeoService_BuildingsByCoordinates_Handler,
		},
		{
			MethodName: "GetAdministrativeDistricts",
			Handler:    _GeoService_GetAdministrativeDistricts_Handler,
		},
		{
			MethodName: "GetDistricts",
			Handler:    _GeoService_GetDistricts_Handler,
		},
		{
			MethodName: "Suggest",
			Handler:    _GeoService_Suggest_Handler,
		},
		{
			MethodName: "GeoLocate",
			Handler:    _GeoService_GeoLocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geo/geo.proto",
}
