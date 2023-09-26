// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.21.9
// source: template_proj/v1/template_proj.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTemplateProjLogin = "/template_proj.v1.TemplateProj/Login"
const OperationTemplateProjTest2 = "/template_proj.v1.TemplateProj/Test2"

type TemplateProjHTTPServer interface {
	// Login 登录
	Login(context.Context, *LoginReq) (*LoginReply, error)
	Test2(context.Context, *Test2Req) (*Test2Reply, error)
}

func RegisterTemplateProjHTTPServer(s *http.Server, srv TemplateProjHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/user/login", _TemplateProj_Login0_HTTP_Handler(srv))
	r.POST("/api/v1/user/test2", _TemplateProj_Test20_HTTP_Handler(srv))
}

func _TemplateProj_Login0_HTTP_Handler(srv TemplateProjHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateProjLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _TemplateProj_Test20_HTTP_Handler(srv TemplateProjHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Test2Req
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTemplateProjTest2)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Test2(ctx, req.(*Test2Req))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Test2Reply)
		return ctx.Result(200, reply)
	}
}

type TemplateProjHTTPClient interface {
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReply, err error)
	Test2(ctx context.Context, req *Test2Req, opts ...http.CallOption) (rsp *Test2Reply, err error)
}

type TemplateProjHTTPClientImpl struct {
	cc *http.Client
}

func NewTemplateProjHTTPClient(client *http.Client) TemplateProjHTTPClient {
	return &TemplateProjHTTPClientImpl{client}
}

func (c *TemplateProjHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/v1/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplateProjLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TemplateProjHTTPClientImpl) Test2(ctx context.Context, in *Test2Req, opts ...http.CallOption) (*Test2Reply, error) {
	var out Test2Reply
	pattern := "/api/v1/user/test2"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTemplateProjTest2))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
