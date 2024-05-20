/*
validate_tx.go

This file contains the logic for validating transactions at various states in the chains consensus.
*/

package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/interface/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (c *Core) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	tx := req.Tx
	logger := c.logger
	var baseCmd commands.Command[any]
	err := c.fromTxBytes(tx, &baseCmd)
	if err != nil {
		return nil, utils.AppError("could not parse a tx as a command", err)
	}

	operation := baseCmd.Operation

	switch operation {
	case commands.Operation{Action: commands.CREATE, Entity: commands.USER}:
		return successCheck()
	case commands.Operation{Action: commands.CREATE, Entity: commands.TRACK}:
		var cmd commands.CreateTrackCommand
		err := c.fromTxBytes(tx, &cmd)
		if err != nil {
			return errorCheck()
		}
		if err := c.checkCreateTrack(&cmd); err != nil {
			logger.Error("invalid track create", err)
			return errorCheck()
		}
	case commands.Operation{Action: commands.CREATE, Entity: commands.FOLLOW}:
	case commands.Operation{Action: commands.CREATE, Entity: commands.REPOST}:
	case commands.Operation{Action: commands.DELETE, Entity: commands.FOLLOW}:
	case commands.Operation{Action: commands.DELETE, Entity: commands.REPOST}:
	default:
		return errorCheck()
	}
	return errorCheck()
}

func successCheck() (*abcitypes.ResponseCheckTx, error) {
	return &abcitypes.ResponseCheckTx{Code: abcitypes.CodeTypeOK}, nil
}

func errorCheck() (*abcitypes.ResponseCheckTx, error) {
	return &abcitypes.ResponseCheckTx{Code: 1}, nil
}
