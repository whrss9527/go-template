// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: template_proj/v1/template_proj.proto

package v1

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

const (
	TemplateProj_Login_FullMethodName = "/template_proj.v1.TemplateProj/Login"
)

// TemplateProjClient is the client API for TemplateProj service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateProjClient interface {
	// 登录
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
}

type templateProjClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateProjClient(cc grpc.ClientConnInterface) TemplateProjClient {
	return &templateProjClient{cc}
}

func (c *templateProjClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, TemplateProj_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateProjServer is the server API for TemplateProj service.
// All implementations must embed UnimplementedTemplateProjServer
// for forward compatibility
type TemplateProjServer interface {
	// 登录
	Login(context.Context, *LoginReq) (*LoginReply, error)
	mustEmbedUnimplementedTemplateProjServer()
}

// UnimplementedTemplateProjServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateProjServer struct {
}

func (UnimplementedTemplateProjServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedTemplateProjServer) mustEmbedUnimplementedTemplateProjServer() {}

// UnsafeTemplateProjServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateProjServer will
// result in compilation errors.
type UnsafeTemplateProjServer interface {
	mustEmbedUnimplementedTemplateProjServer()
}

func RegisterTemplateProjServer(s grpc.ServiceRegistrar, srv TemplateProjServer) {
	s.RegisterService(&TemplateProj_ServiceDesc, srv)
}

func _TemplateProj_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateProjServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateProj_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateProjServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateProj_ServiceDesc is the grpc.ServiceDesc for TemplateProj service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateProj_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template_proj.v1.TemplateProj",
	HandlerType: (*TemplateProjServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TemplateProj_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template_proj/v1/template_proj.proto",
}
