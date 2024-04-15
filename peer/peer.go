package peer

import (
	"context"
	"time"

	"github.com/alecsavvy/clockwise/server/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Peer struct {
	conn     *grpc.ClientConn
	rpc      gen.NodeServiceClient
	endpoint string
}

func NewPeer(endpoint string) (*Peer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             2 * time.Second,
			PermitWithoutStream: true,
		}))
	if err != nil {
		return nil, err
	}

	node := gen.NewNodeServiceClient(conn)

	return &Peer{
		conn:     conn,
		endpoint: endpoint,
		rpc:      node,
	}, nil
}
