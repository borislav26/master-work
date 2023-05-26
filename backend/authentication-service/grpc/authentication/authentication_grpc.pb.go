// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: authentication.proto

package authentication_grpc

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

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	CheckUserAuthentication(ctx context.Context, in *CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*CheckUserAuthenticationResponse, error)
	GetUserEmailFromToken(ctx context.Context, in *CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*GetUserEmailFromTokenResponse, error)
	FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserByEmailResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) CheckUserAuthentication(ctx context.Context, in *CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*CheckUserAuthenticationResponse, error) {
	out := new(CheckUserAuthenticationResponse)
	err := c.cc.Invoke(ctx, "/authentication.AuthenticationService/CheckUserAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetUserEmailFromToken(ctx context.Context, in *CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*GetUserEmailFromTokenResponse, error) {
	out := new(GetUserEmailFromTokenResponse)
	err := c.cc.Invoke(ctx, "/authentication.AuthenticationService/GetUserEmailFromToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) FindUserByEmail(ctx context.Context, in *FindUserByEmailRequest, opts ...grpc.CallOption) (*FindUserByEmailResponse, error) {
	out := new(FindUserByEmailResponse)
	err := c.cc.Invoke(ctx, "/authentication.AuthenticationService/FindUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	CheckUserAuthentication(context.Context, *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error)
	GetUserEmailFromToken(context.Context, *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error)
	FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserByEmailResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) CheckUserAuthentication(context.Context, *CheckUserAuthenticationRequest) (*CheckUserAuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAuthentication not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetUserEmailFromToken(context.Context, *CheckUserAuthenticationRequest) (*GetUserEmailFromTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEmailFromToken not implemented")
}
func (UnimplementedAuthenticationServiceServer) FindUserByEmail(context.Context, *FindUserByEmailRequest) (*FindUserByEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmail not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_CheckUserAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).CheckUserAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.AuthenticationService/CheckUserAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).CheckUserAuthentication(ctx, req.(*CheckUserAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetUserEmailFromToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetUserEmailFromToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.AuthenticationService/GetUserEmailFromToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetUserEmailFromToken(ctx, req.(*CheckUserAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_FindUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).FindUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.AuthenticationService/FindUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).FindUserByEmail(ctx, req.(*FindUserByEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUserAuthentication",
			Handler:    _AuthenticationService_CheckUserAuthentication_Handler,
		},
		{
			MethodName: "GetUserEmailFromToken",
			Handler:    _AuthenticationService_GetUserEmailFromToken_Handler,
		},
		{
			MethodName: "FindUserByEmail",
			Handler:    _AuthenticationService_FindUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
