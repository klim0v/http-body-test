package service

import (
	"bytes"
	"encoding/csv"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/klim0v/http-body-test/pb"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
}

func (*Service) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*pb.PongResponse, error) {
	return &pb.PongResponse{Pong: "OK"}, nil
}

func (*Service) GetResource(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	writer.Comma = ';'

	err := writer.Write([]string{"1", "gopher conference", "$100"})
	if err != nil {
		return new(httpbody.HttpBody), status.New(codes.Internal, err.Error()).Err()
	}

	writer.Flush()

	response := &httpbody.HttpBody{
		ContentType: "application/octet-stream",
		Data:        buf.Bytes(),
	}
	return response, nil
}

func NewService() *Service {
	return &Service{}
}
