package service

import (
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
)

type Service struct {
}

func (*Service) GetResource(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	response := &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello World"),
	}
	return response, nil
}

func NewService() *Service {
	return &Service{}
}
