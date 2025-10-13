package core

import (
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/rpc/client/local"
)

type Core struct {
	node *node.Node
	rpc  *local.Local
}
