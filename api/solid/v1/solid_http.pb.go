// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v3.19.4
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
const OperationSolidServiceDeleteSolid = "/api.solid.v1.SolidService/DeleteSolid"
const OperationSolidServiceGetSolid = "/api.solid.v1.SolidService/GetSolid"
const OperationSolidServiceListSolid = "/api.solid.v1.SolidService/ListSolid"
const OperationSolidServiceUpdateSolid = "/api.solid.v1.SolidService/UpdateSolid"

type SolidServiceHTTPServer interface {
	CreateSolid(context.Context, *CreateSolidRequest) (*CreateSolidReply, error)
	DeleteSolid(context.Context, *DeleteSolidRequest) (*DeleteSolidReply, error)
	GetSolid(context.Context, *GetSolidRequest) (*GetSolidReply, error)
	ListSolid(context.Context, *ListSolidRequest) (*ListSolidReply, error)
	UpdateSolid(context.Context, *UpdateSolidRequest) (*UpdateSolidReply, error)
}

func RegisterSolidServiceHTTPServer(s *http.Server, srv SolidServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/solid", _SolidService_CreateSolid0_HTTP_Handler(srv))
	r.PUT("/v1/solid/{id}", _SolidService_UpdateSolid0_HTTP_Handler(srv))
	r.DELETE("/v1/solid/{id}", _SolidService_DeleteSolid0_HTTP_Handler(srv))
	r.GET("/v1/solid/{id}", _SolidService_GetSolid0_HTTP_Handler(srv))
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

func _SolidService_UpdateSolid0_HTTP_Handler(srv SolidServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSolidRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolidServiceUpdateSolid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSolid(ctx, req.(*UpdateSolidRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSolidReply)
		return ctx.Result(200, reply)
	}
}

func _SolidService_DeleteSolid0_HTTP_Handler(srv SolidServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSolidRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolidServiceDeleteSolid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSolid(ctx, req.(*DeleteSolidRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSolidReply)
		return ctx.Result(200, reply)
	}
}

func _SolidService_GetSolid0_HTTP_Handler(srv SolidServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSolidRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSolidServiceGetSolid)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSolid(ctx, req.(*GetSolidRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSolidReply)
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
	DeleteSolid(ctx context.Context, req *DeleteSolidRequest, opts ...http.CallOption) (rsp *DeleteSolidReply, err error)
	GetSolid(ctx context.Context, req *GetSolidRequest, opts ...http.CallOption) (rsp *GetSolidReply, err error)
	ListSolid(ctx context.Context, req *ListSolidRequest, opts ...http.CallOption) (rsp *ListSolidReply, err error)
	UpdateSolid(ctx context.Context, req *UpdateSolidRequest, opts ...http.CallOption) (rsp *UpdateSolidReply, err error)
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

func (c *SolidServiceHTTPClientImpl) DeleteSolid(ctx context.Context, in *DeleteSolidRequest, opts ...http.CallOption) (*DeleteSolidReply, error) {
	var out DeleteSolidReply
	pattern := "/v1/solid/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSolidServiceDeleteSolid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SolidServiceHTTPClientImpl) GetSolid(ctx context.Context, in *GetSolidRequest, opts ...http.CallOption) (*GetSolidReply, error) {
	var out GetSolidReply
	pattern := "/v1/solid/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSolidServiceGetSolid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
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

func (c *SolidServiceHTTPClientImpl) UpdateSolid(ctx context.Context, in *UpdateSolidRequest, opts ...http.CallOption) (*UpdateSolidReply, error) {
	var out UpdateSolidReply
	pattern := "/v1/solid/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSolidServiceUpdateSolid))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}