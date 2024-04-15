package server

import (
	"context"
	"time"

	pb "github.com/alecsavvy/clockwise/server/gen"
)

type ServerImpl struct {
	pb.UnimplementedNodeServiceServer
}

func (s *ServerImpl) GetNodeHealth(ctx context.Context, in *pb.NodeHealthRequest) (*pb.NodeHealthResponse, error) {
	currentUTC := time.Now().UTC().Unix()
	return &pb.NodeHealthResponse{Healthy: true, Timestamp: currentUTC}, nil
}
