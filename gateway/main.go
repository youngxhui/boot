package main

import (
	wg "boot/protos"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/youngxhui/power/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc服务地址
	endpoint := "127.0.0.1:9091"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc

	err := wg.RegisterUserSrvHandlerFromEndpoint(ctx, mux, endpoint, opts)
	err = wg.RegisterPingHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		grpclog.Fatalf("Register handler err:%v\n", err)
	}

	log.Info("HTTP Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
