/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"
	"fmt"

	"github.com/alecsavvy/clockwise/core/chain/handlers"
	"github.com/alecsavvy/clockwise/core/db"
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
		// todo: unique id?
		ID:          fmt.Sprintf("%d", rfb.Height),
		Blocknumber: int32(rfb.Height),
		Blockhash:   string(rfb.Hash),
	})

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
