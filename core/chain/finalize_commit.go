/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"

	"github.com/alecsavvy/clockwise/core/chain/handlers"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {

	dbTx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	qtx := a.queries.WithTx(dbTx)

	results, err := handlers.RootHandler(qtx, rfb.Txs)

	a.currentTx = dbTx

	return &abcitypes.ResponseFinalizeBlock{TxResults: results}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (a *Application) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	if a.currentTx != nil {
		err := a.currentTx.Commit(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &abcitypes.ResponseCommit{}, nil
}
