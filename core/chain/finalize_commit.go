/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"

	"github.com/alecsavvy/clockwise/core/chain/handlers"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {

	dbTx, err := a.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	qtx := a.queries.WithTx(dbTx)

	// insert block as first tx
	err = qtx.CreateBlock(ctx, db.CreateBlockParams{
		Blocknumber: int32(rfb.Height),
		Blockhash:   rfb.Hash,
	})
	if err != nil {
		return nil, utils.AppError("error inserting current block", err)
	}

	results, err := handlers.RootHandler(qtx, int32(rfb.Height), rfb.Txs)
	if err != nil {
		return nil, utils.AppError("error in root handler", err)
	}

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
