/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func (c *Core) rootHandler(qtx *db.Queries, rfb *abcitypes.RequestFinalizeBlock) ([]*abcitypes.ExecTxResult, error) {
	createdAt := rfb.Height
	txs := rfb.Txs
	var txResults = make([]*abcitypes.ExecTxResult, len(txs))

	if len(txs) <= 0 {
		return txResults, nil
	}

	// init tx results
	for i, tx := range txs {
		var baseCmd commands.Command[any]
		err := c.fromTxBytes(tx, &baseCmd)
		if err != nil {
			return nil, utils.AppError("could not parse a tx as a command", err)
		}

		operation := baseCmd.Operation

		switch operation {
		case commands.Operation{Action: commands.CREATE, Entity: commands.USER}:
			txResult, err := c.handleCreateUser(qtx, createdAt, tx)
			if err != nil {
				return nil, utils.AppError("cannot handle create user", err)
			}
			txResults[i] = txResult
		case commands.Operation{Action: commands.CREATE, Entity: commands.TRACK}:
			txResult, err := c.handleCreateTrack(qtx, createdAt, tx)
			if err != nil {
				return nil, utils.AppError("cannot handle create track", err)
			}
			txResults[i] = txResult
		case commands.Operation{Action: commands.CREATE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.CREATE, Entity: commands.REPOST}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.REPOST}:
		default:
			return nil, utils.AppError("found transaction with invalid operation", nil)
		}
	}
	return txResults, nil
}

// Prepares a block for commitment and provides final validation
func (c *Core) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {

	dbTx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	qtx := c.db.WithTx(dbTx)

	// insert block as first tx
	err = qtx.CreateBlock(ctx, db.CreateBlockParams{
		Blocknumber: rfb.Height,
		Blockhash:   rfb.Hash,
	})
	if err != nil {
		return nil, utils.AppError("error inserting current block", err)
	}

	results, err := c.rootHandler(qtx, rfb)
	if err != nil {
		return nil, utils.AppError("error in root handler", err)
	}

	c.currentTx = dbTx

	return &abcitypes.ResponseFinalizeBlock{TxResults: results}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (c *Core) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	if c.currentTx != nil {
		err := c.currentTx.Commit(ctx)
		if err != nil {
			return nil, err
		}
	}
	c.currentTx = nil
	return &abcitypes.ResponseCommit{}, nil
}
