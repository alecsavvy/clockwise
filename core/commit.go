/*
finalize_commit.go

This file contains both the finalize and the commit step in the consensus algorithm.
*/
package core

import (
	"context"
	"fmt"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Prepares a block for commitment and provides final validation
func (c *Core) FinalizeBlock(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) (*abcitypes.ResponseFinalizeBlock, error) {
	var txResults = make([]*abcitypes.ExecTxResult, len(rfb.Txs))
	for i, tx := range rfb.Txs {
		var me ManageEntity
		c.fromTxBytes(tx, &me)
		txResults[i] = &abcitypes.ExecTxResult{
			Code: abcitypes.CodeTypeOK,
			Events: []abcitypes.Event{
				{
					Type: fmt.Sprintf("%s%s", me.EntityType, me.Action),
					Attributes: []abcitypes.EventAttribute{
						{Key: "requestId", Value: me.RequestID},
					},
				},
			},
		}
	}
	return &abcitypes.ResponseFinalizeBlock{TxResults: txResults}, nil
}

// Writes the state changes to the database after checking and finalizing a block
func (c *Core) Commit(ctx context.Context, req *abcitypes.RequestCommit) (*abcitypes.ResponseCommit, error) {
	// persist anything if we did
	return &abcitypes.ResponseCommit{}, nil
}
