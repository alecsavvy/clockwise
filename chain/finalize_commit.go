/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	dbTx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer dbTx.Rollback(ctx)

	_ = a.queries.WithTx(dbTx)

	// set transaction to be committed in commit step
	a.currentTx = dbTx
	return &abcitypes.ResponseFinalizeBlock{}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (a *Application) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	// check transactions again
	a.currentTx.Commit(ctx)
	return &abcitypes.ResponseCommit{}, nil
}
