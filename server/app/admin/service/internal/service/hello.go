package service

import (
	"context"
	helloV1 "mosong/api/gen/go/hello/service/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) Hello(ctx context.Context, _ *emptypb.Empty) (*helloV1.HelloResponse, error) {
	return &helloV1.HelloResponse{
		Hello: "world",
	}, nil
}
