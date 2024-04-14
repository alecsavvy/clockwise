package grpc

import (
	"context"
	"time"

	pb "github.com/alecsavvy/clockwise/grpc/gen"
)

type ServerImpl struct {
	pb.UnimplementedNodeServiceServer
}

func (s *ServerImpl) SayHello(ctx context.Context, in *pb.NodeHealthRequest) (*pb.NodeHealthResponse, error) {
	currentUTC := time.Now().UTC().Unix()
	return &pb.NodeHealthResponse{Healthy: true, Timestamp: currentUTC}, nil
}
