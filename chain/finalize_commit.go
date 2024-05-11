/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"

	"github.com/alecsavvy/clockwise/db"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/jackc/pgx/v5/pgtype"
)

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	dbTx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer dbTx.Rollback(ctx)

	qtx := a.queries.WithTx(dbTx)

	// do db logic with qtx
	blockTime := pgtype.Date{
		Time:  rfb.Time,
		Valid: true,
	}
	qtx.InsertBlock(ctx, db.InsertBlockParams{
		Blocknumber: rfb.Height,
		Blocktime:   blockTime,
	})

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
