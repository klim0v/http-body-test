package main

import (
	"flag"
	"github.com/klim0v/http-body-test/service"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"

	gw "github.com/klim0v/http-body-test/pb"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.SetHTTPBodyMarshaler)
	err := gw.RegisterServiceHandlerClient(ctx, mux, service.NewService())
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
