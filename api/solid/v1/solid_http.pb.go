// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.21.8
// source: solid/v1/solid.proto

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

const OperationSolidServiceCreateSolid = "/api.solid.v1.SolidService/CreateSolid"
const OperationSolidServiceListSolid = "/api.solid.v1.SolidService/ListSolid"

type SolidServiceHTTPServer interface {
	CreateSolid(context.Context, *CreateSolidRequest) (*CreateSolidReply, error)
	ListSolid(context.Context, *ListSolidRequest) (*ListSolidReply, error)
}

func RegisterSolidServiceHTTPServer(s *http.Server, srv SolidServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/solid", _SolidService_CreateSolid0_HTTP_Handler(srv))
	r.GET("/v1/solid", _SolidService_ListSolid0_HTTP_Handler(srv))
}

func _SolidService_CreateSolid0_HTTP_Handler(srv SolidServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSolidRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolidServiceCreateSolid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSolid(ctx, req.(*CreateSolidRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSolidReply)
		return ctx.Result(200, reply)
	}
}

func _SolidService_ListSolid0_HTTP_Handler(srv SolidServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSolidRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolidServiceListSolid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSolid(ctx, req.(*ListSolidRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSolidReply)
		return ctx.Result(200, reply)
	}
}

type SolidServiceHTTPClient interface {
	CreateSolid(ctx context.Context, req *CreateSolidRequest, opts ...http.CallOption) (rsp *CreateSolidReply, err error)
	ListSolid(ctx context.Context, req *ListSolidRequest, opts ...http.CallOption) (rsp *ListSolidReply, err error)
}

type SolidServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSolidServiceHTTPClient(client *http.Client) SolidServiceHTTPClient {
	return &SolidServiceHTTPClientImpl{client}
}

func (c *SolidServiceHTTPClientImpl) CreateSolid(ctx context.Context, in *CreateSolidRequest, opts ...http.CallOption) (*CreateSolidReply, error) {
	var out CreateSolidReply
	pattern := "/v1/solid"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSolidServiceCreateSolid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SolidServiceHTTPClientImpl) ListSolid(ctx context.Context, in *ListSolidRequest, opts ...http.CallOption) (*ListSolidReply, error) {
	var out ListSolidReply
	pattern := "/v1/solid"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSolidServiceListSolid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
