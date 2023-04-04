// Package service implements the test gRPC service
package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/DolceTriade/monorepo/proto"
)

type Service struct {
	log *zap.SugaredLogger
}

func Register(log *zap.SugaredLogger, s *grpc.Server) {
	svc := &Service{
		log: log,
	}
	pb.RegisterTestServiceServer(s, svc)
}

func (s *Service) Test(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Response: req.Request,
	}, nil
}
