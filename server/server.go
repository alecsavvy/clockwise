package server

import (
	"fmt"
	"net"

	"github.com/alecsavvy/clockwise/common"
	pb "github.com/alecsavvy/clockwise/server/gen"
	"google.golang.org/grpc"
)

type RpcService struct {
	impl   *grpc.Server
	lis    net.Listener
	config *common.Config
}

func New(config *common.Config) (*RpcService, error) {
	lis, err := net.Listen("tcp", config.NodeEndpoint)
	if err != nil {
		return nil, err
	}

	impl := grpc.NewServer()
	pb.RegisterNodeServiceServer(impl, &ServerImpl{})

	return &RpcService{
		impl:   impl,
		lis:    lis,
		config: config,
	}, nil
}

func (server *RpcService) Run() error {
	fmt.Printf("grpc server starting up on %s\n", server.config.NodeEndpoint)
	return server.impl.Serve(server.lis)
}
