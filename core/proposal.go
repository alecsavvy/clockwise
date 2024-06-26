/*
proposal.go

Contains the block proposal logic. Creates and receives new transactions from the network.
*/
package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a new block proposal for the network
func (c *Core) PrepareProposal(_ context.Context, proposal *abcitypes.RequestPrepareProposal) (*abcitypes.ResponsePrepareProposal, error) {
	// TODO: reorder transactions in here
	return &abcitypes.ResponsePrepareProposal{Txs: proposal.Txs}, nil
}

// Processes block proposal from the network created by PrepareProposal
func (c *Core) ProcessProposal(context.Context, *abcitypes.RequestProcessProposal) (*abcitypes.ResponseProcessProposal, error) {
	return &abcitypes.ResponseProcessProposal{Status: abcitypes.ResponseProcessProposal_ACCEPT}, nil
}
