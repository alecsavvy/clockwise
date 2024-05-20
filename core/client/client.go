/* the client that external modules use like grpc and clis */
package client

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/rpc/client/local"
)

type Core struct {
	logger *utils.Logger
	node   *node.Node
	rpc    *local.Local
	db     *db.Queries
}

func NewCore(logger *utils.Logger, node *node.Node, db *db.Queries) *Core {
	rpc := local.New(node)
	return &Core{
		logger: logger,
		node:   node,
		rpc:    rpc,
		db:     db,
	}
}

func (c *Core) Run() error {
	return nil
}
