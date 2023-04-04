// Package main runs the test grpc server
package main

import (
	"flag"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/DolceTriade/monorepo/go/server/service"
)

var port = flag.Int("port", 0, "grpc port")

func main() {
	zlog := zap.Must(zap.NewDevelopment())
	defer zlog.Sync()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		zlog.Fatal("failed to listen: %v", zap.Error(err))
	}
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	service.Register(zlog.Sugar(), s)
	zlog.Info("server listening", zap.String("addr", lis.Addr().String()))
	if err := s.Serve(lis); err != nil {
		zlog.Fatal("failed to serve", zap.Error(err))
	}
}
