/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package chain

import (
	"context"

	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a block for commitment and provides final validation
func (a *Application) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	var txs = make([]*abcitypes.ExecTxResult, len(rfb.Txs))

	if len(rfb.Txs) > 0 {
		// only doing on tx per block rn
		tx, err := chainutils.FromTxBytes[commands.CreateUserCommand](rfb.Txs[0])
		if err != nil {
			return nil, utils.AppError("could not decode tx in finalize", err)
		}

		dbTx, err := a.pool.Begin(ctx)
		if err != nil {
			return nil, err
		}

		qtx := a.queries.WithTx(dbTx)

		params := &db.CreateUserParams{
			ID:      tx.ID,
			Handle:  tx.Handle,
			Bio:     tx.Bio,
			Address: tx.Address,
		}
		err = qtx.CreateUser(ctx, *params)
		if err != nil {
			return nil, err
		}

		// set transaction to be committed in commit step
		a.currentTx = dbTx

		txs[0] = &abcitypes.ExecTxResult{
			Code: 0,
		}
	}
	return &abcitypes.ResponseFinalizeBlock{TxResults: txs}, nil
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
