// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: solid/v1/solid.proto

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

// SolidServiceClient is the client API for SolidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolidServiceClient interface {
	CreateSolid(ctx context.Context, in *CreateSolidRequest, opts ...grpc.CallOption) (*CreateSolidReply, error)
	//	rpc UpdateSolid (UpdateSolidRequest) returns (UpdateSolidReply) {
	//		option (google.api.http) = {
	//			put: "/v1/solid/{id}"
	//			body: "*"
	//		};
	//	};
	//
	//	rpc DeleteSolid (DeleteSolidRequest) returns (DeleteSolidReply) {
	//		option (google.api.http) = {
	//			delete: "/v1/solid/{id}"
	//		};
	//	};
	//
	//	rpc GetSolid (GetSolidRequest) returns (GetSolidReply) {
	//		option (google.api.http) = {
	//			get: "/v1/solid/{id}"
	//		};
	//	};
	ListSolid(ctx context.Context, in *ListSolidRequest, opts ...grpc.CallOption) (*ListSolidReply, error)
}

type solidServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSolidServiceClient(cc grpc.ClientConnInterface) SolidServiceClient {
	return &solidServiceClient{cc}
}

func (c *solidServiceClient) CreateSolid(ctx context.Context, in *CreateSolidRequest, opts ...grpc.CallOption) (*CreateSolidReply, error) {
	out := new(CreateSolidReply)
	err := c.cc.Invoke(ctx, "/api.solid.v1.SolidService/CreateSolid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *solidServiceClient) ListSolid(ctx context.Context, in *ListSolidRequest, opts ...grpc.CallOption) (*ListSolidReply, error) {
	out := new(ListSolidReply)
	err := c.cc.Invoke(ctx, "/api.solid.v1.SolidService/ListSolid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SolidServiceServer is the server API for SolidService service.
// All implementations must embed UnimplementedSolidServiceServer
// for forward compatibility
type SolidServiceServer interface {
	CreateSolid(context.Context, *CreateSolidRequest) (*CreateSolidReply, error)
	//	rpc UpdateSolid (UpdateSolidRequest) returns (UpdateSolidReply) {
	//		option (google.api.http) = {
	//			put: "/v1/solid/{id}"
	//			body: "*"
	//		};
	//	};
	//
	//	rpc DeleteSolid (DeleteSolidRequest) returns (DeleteSolidReply) {
	//		option (google.api.http) = {
	//			delete: "/v1/solid/{id}"
	//		};
	//	};
	//
	//	rpc GetSolid (GetSolidRequest) returns (GetSolidReply) {
	//		option (google.api.http) = {
	//			get: "/v1/solid/{id}"
	//		};
	//	};
	ListSolid(context.Context, *ListSolidRequest) (*ListSolidReply, error)
	mustEmbedUnimplementedSolidServiceServer()
}

// UnimplementedSolidServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSolidServiceServer struct {
}

func (UnimplementedSolidServiceServer) CreateSolid(context.Context, *CreateSolidRequest) (*CreateSolidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSolid not implemented")
}
func (UnimplementedSolidServiceServer) ListSolid(context.Context, *ListSolidRequest) (*ListSolidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSolid not implemented")
}
func (UnimplementedSolidServiceServer) mustEmbedUnimplementedSolidServiceServer() {}

// UnsafeSolidServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolidServiceServer will
// result in compilation errors.
type UnsafeSolidServiceServer interface {
	mustEmbedUnimplementedSolidServiceServer()
}

func RegisterSolidServiceServer(s grpc.ServiceRegistrar, srv SolidServiceServer) {
	s.RegisterService(&SolidService_ServiceDesc, srv)
}

func _SolidService_CreateSolid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSolidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolidServiceServer).CreateSolid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.solid.v1.SolidService/CreateSolid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolidServiceServer).CreateSolid(ctx, req.(*CreateSolidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SolidService_ListSolid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSolidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SolidServiceServer).ListSolid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.solid.v1.SolidService/ListSolid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SolidServiceServer).ListSolid(ctx, req.(*ListSolidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SolidService_ServiceDesc is the grpc.ServiceDesc for SolidService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolidService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.solid.v1.SolidService",
	HandlerType: (*SolidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSolid",
			Handler:    _SolidService_CreateSolid_Handler,
		},
		{
			MethodName: "ListSolid",
			Handler:    _SolidService_ListSolid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solid/v1/solid.proto",
}
