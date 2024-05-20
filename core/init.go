/*
init.go

Contains chain initialization logic. Determines whether the chain is new, the sync state of a restarted node, etc.
*/
package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Info implements types.Application.
func (c *Core) Info(context.Context, *abcitypes.RequestInfo) (*abcitypes.ResponseInfo, error) {
	return &abcitypes.ResponseInfo{}, nil
}

// InitChain implements types.Application.
func (c *Core) InitChain(context.Context, *abcitypes.RequestInitChain) (*abcitypes.ResponseInitChain, error) {
	return &abcitypes.ResponseInitChain{}, nil
}
