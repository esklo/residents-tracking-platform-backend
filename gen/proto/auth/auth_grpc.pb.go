// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: auth/auth.proto

package auth

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
	AuthService_Login_FullMethodName                       = "/auth.AuthService/Login"
	AuthService_Logout_FullMethodName                      = "/auth.AuthService/Logout"
	AuthService_RequestPublicKeyAttestation_FullMethodName = "/auth.AuthService/RequestPublicKeyAttestation"
	AuthService_PublicKeyAttestation_FullMethodName        = "/auth.AuthService/PublicKeyAttestation"
	AuthService_RequestPublicKeyAssertion_FullMethodName   = "/auth.AuthService/RequestPublicKeyAssertion"
	AuthService_PublicKeyAssertion_FullMethodName          = "/auth.AuthService/PublicKeyAssertion"
	AuthService_GetPublicKeys_FullMethodName               = "/auth.AuthService/GetPublicKeys"
	AuthService_DeletePublicKey_FullMethodName             = "/auth.AuthService/DeletePublicKey"
	AuthService_ChangePassword_FullMethodName              = "/auth.AuthService/ChangePassword"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	RequestPublicKeyAttestation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PublicKeyCredentialsResponse, error)
	PublicKeyAttestation(ctx context.Context, in *PublicKeyCredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RequestPublicKeyAssertion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PublicKeyCredentialsResponse, error)
	PublicKeyAssertion(ctx context.Context, in *PublicKeyCredentialsRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetPublicKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPublicKeysResponse, error)
	DeletePublicKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RequestPublicKeyAttestation(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PublicKeyCredentialsResponse, error) {
	out := new(PublicKeyCredentialsResponse)
	err := c.cc.Invoke(ctx, AuthService_RequestPublicKeyAttestation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PublicKeyAttestation(ctx context.Context, in *PublicKeyCredentialsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AuthService_PublicKeyAttestation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RequestPublicKeyAssertion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PublicKeyCredentialsResponse, error) {
	out := new(PublicKeyCredentialsResponse)
	err := c.cc.Invoke(ctx, AuthService_RequestPublicKeyAssertion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PublicKeyAssertion(ctx context.Context, in *PublicKeyCredentialsRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthService_PublicKeyAssertion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetPublicKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetPublicKeysResponse, error) {
	out := new(GetPublicKeysResponse)
	err := c.cc.Invoke(ctx, AuthService_GetPublicKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeletePublicKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AuthService_DeletePublicKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AuthService_ChangePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *empty.Empty) (*empty.Empty, error)
	RequestPublicKeyAttestation(context.Context, *empty.Empty) (*PublicKeyCredentialsResponse, error)
	PublicKeyAttestation(context.Context, *PublicKeyCredentialsRequest) (*empty.Empty, error)
	RequestPublicKeyAssertion(context.Context, *empty.Empty) (*PublicKeyCredentialsResponse, error)
	PublicKeyAssertion(context.Context, *PublicKeyCredentialsRequest) (*LoginResponse, error)
	GetPublicKeys(context.Context, *empty.Empty) (*GetPublicKeysResponse, error)
	DeletePublicKey(context.Context, *Key) (*empty.Empty, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) RequestPublicKeyAttestation(context.Context, *empty.Empty) (*PublicKeyCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPublicKeyAttestation not implemented")
}
func (UnimplementedAuthServiceServer) PublicKeyAttestation(context.Context, *PublicKeyCredentialsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKeyAttestation not implemented")
}
func (UnimplementedAuthServiceServer) RequestPublicKeyAssertion(context.Context, *empty.Empty) (*PublicKeyCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPublicKeyAssertion not implemented")
}
func (UnimplementedAuthServiceServer) PublicKeyAssertion(context.Context, *PublicKeyCredentialsRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicKeyAssertion not implemented")
}
func (UnimplementedAuthServiceServer) GetPublicKeys(context.Context, *empty.Empty) (*GetPublicKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKeys not implemented")
}
func (UnimplementedAuthServiceServer) DeletePublicKey(context.Context, *Key) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublicKey not implemented")
}
func (UnimplementedAuthServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RequestPublicKeyAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RequestPublicKeyAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RequestPublicKeyAttestation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RequestPublicKeyAttestation(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PublicKeyAttestation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKeyCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PublicKeyAttestation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PublicKeyAttestation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PublicKeyAttestation(ctx, req.(*PublicKeyCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RequestPublicKeyAssertion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RequestPublicKeyAssertion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RequestPublicKeyAssertion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RequestPublicKeyAssertion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PublicKeyAssertion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKeyCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PublicKeyAssertion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_PublicKeyAssertion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PublicKeyAssertion(ctx, req.(*PublicKeyCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetPublicKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetPublicKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetPublicKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetPublicKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeletePublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeletePublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeletePublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeletePublicKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "RequestPublicKeyAttestation",
			Handler:    _AuthService_RequestPublicKeyAttestation_Handler,
		},
		{
			MethodName: "PublicKeyAttestation",
			Handler:    _AuthService_PublicKeyAttestation_Handler,
		},
		{
			MethodName: "RequestPublicKeyAssertion",
			Handler:    _AuthService_RequestPublicKeyAssertion_Handler,
		},
		{
			MethodName: "PublicKeyAssertion",
			Handler:    _AuthService_PublicKeyAssertion_Handler,
		},
		{
			MethodName: "GetPublicKeys",
			Handler:    _AuthService_GetPublicKeys_Handler,
		},
		{
			MethodName: "DeletePublicKey",
			Handler:    _AuthService_DeletePublicKey_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
