/*
init.go

Contains chain initialization logic. Determines whether the chain is new, the sync state of a restarted node, etc.
*/
package chain

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Info implements types.Application.
func (a *Application) Info(context.Context, *abcitypes.RequestInfo) (*abcitypes.ResponseInfo, error) {
	return &abcitypes.ResponseInfo{}, nil
}

// InitChain implements types.Application.
func (a *Application) InitChain(context.Context, *abcitypes.RequestInitChain) (*abcitypes.ResponseInitChain, error) {
	return &abcitypes.ResponseInitChain{}, nil
}
