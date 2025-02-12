// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: url_shortener/url_shortener.proto

package url_shortener_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	URLService_SaveURL_FullMethodName  = "/url_shortener.URLService/SaveURL"
	URLService_FetchURL_FullMethodName = "/url_shortener.URLService/FetchURL"
)

// URLServiceClient is the client API for URLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLServiceClient interface {
	SaveURL(ctx context.Context, in *SaveURLRequest, opts ...grpc.CallOption) (*SaveURLResponse, error)
	FetchURL(ctx context.Context, in *FetchURLRequest, opts ...grpc.CallOption) (*FetchURLResponse, error)
}

type uRLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewURLServiceClient(cc grpc.ClientConnInterface) URLServiceClient {
	return &uRLServiceClient{cc}
}

func (c *uRLServiceClient) SaveURL(ctx context.Context, in *SaveURLRequest, opts ...grpc.CallOption) (*SaveURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveURLResponse)
	err := c.cc.Invoke(ctx, URLService_SaveURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLServiceClient) FetchURL(ctx context.Context, in *FetchURLRequest, opts ...grpc.CallOption) (*FetchURLResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchURLResponse)
	err := c.cc.Invoke(ctx, URLService_FetchURL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLServiceServer is the server API for URLService service.
// All implementations must embed UnimplementedURLServiceServer
// for forward compatibility.
type URLServiceServer interface {
	SaveURL(context.Context, *SaveURLRequest) (*SaveURLResponse, error)
	FetchURL(context.Context, *FetchURLRequest) (*FetchURLResponse, error)
	mustEmbedUnimplementedURLServiceServer()
}

// UnimplementedURLServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedURLServiceServer struct{}

func (UnimplementedURLServiceServer) SaveURL(context.Context, *SaveURLRequest) (*SaveURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveURL not implemented")
}
func (UnimplementedURLServiceServer) FetchURL(context.Context, *FetchURLRequest) (*FetchURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchURL not implemented")
}
func (UnimplementedURLServiceServer) mustEmbedUnimplementedURLServiceServer() {}
func (UnimplementedURLServiceServer) testEmbeddedByValue()                    {}

// UnsafeURLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLServiceServer will
// result in compilation errors.
type UnsafeURLServiceServer interface {
	mustEmbedUnimplementedURLServiceServer()
}

func RegisterURLServiceServer(s grpc.ServiceRegistrar, srv URLServiceServer) {
	// If the following call pancis, it indicates UnimplementedURLServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&URLService_ServiceDesc, srv)
}

func _URLService_SaveURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLServiceServer).SaveURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLService_SaveURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLServiceServer).SaveURL(ctx, req.(*SaveURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLService_FetchURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLServiceServer).FetchURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLService_FetchURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLServiceServer).FetchURL(ctx, req.(*FetchURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLService_ServiceDesc is the grpc.ServiceDesc for URLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "url_shortener.URLService",
	HandlerType: (*URLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveURL",
			Handler:    _URLService_SaveURL_Handler,
		},
		{
			MethodName: "FetchURL",
			Handler:    _URLService_FetchURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "url_shortener/url_shortener.proto",
}
