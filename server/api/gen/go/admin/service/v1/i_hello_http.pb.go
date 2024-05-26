// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: admin/service/v1/i_hello.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "mosong/api/gen/go/hello/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationHelloServiceHello = "/admin.service.v1.HelloService/Hello"

type HelloServiceHTTPServer interface {
	Hello(context.Context, *emptypb.Empty) (*v1.HelloResponse, error)
}

func RegisterHelloServiceHTTPServer(s *http.Server, srv HelloServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/hellos", _HelloService_Hello0_HTTP_Handler(srv))
}

func _HelloService_Hello0_HTTP_Handler(srv HelloServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationHelloServiceHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Hello(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.HelloResponse)
		return ctx.Result(200, reply)
	}
}

type HelloServiceHTTPClient interface {
	Hello(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *v1.HelloResponse, err error)
}

type HelloServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewHelloServiceHTTPClient(client *http.Client) HelloServiceHTTPClient {
	return &HelloServiceHTTPClientImpl{client}
}

func (c *HelloServiceHTTPClientImpl) Hello(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*v1.HelloResponse, error) {
	var out v1.HelloResponse
	pattern := "/admin/v1/hellos"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationHelloServiceHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
