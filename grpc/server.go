package grpc

import (
	"fmt"
	"net"

	pb "github.com/alecsavvy/clockwise/grpc/gen"
	"google.golang.org/grpc"
)

type Server struct {
	impl *grpc.Server
	lis  net.Listener
}

func New(host string) (*Server, error) {
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, err
	}

	impl := grpc.NewServer()
	pb.RegisterNodeServiceServer(impl, &ServerImpl{})

	return &Server{
		impl: impl,
		lis:  lis,
	}, nil
}

func (server *Server) Serve() error {
	fmt.Println("grpc server starting up...")
	return server.impl.Serve(server.lis)
}
