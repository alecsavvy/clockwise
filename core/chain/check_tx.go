/*
validate_tx.go

This file contains the logic for validating transactions at various states in the chains consensus.
*/

package chain

import (
	"context"

	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (a *Application) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	tx, err := chainutils.FromTxBytes[commands.CreateUserCommand](req.Tx)
	if err != nil {
		return nil, utils.AppError("could not decode tx", err)
	}
	a.logger.Info("checking tx", "tx", tx)
	return &abcitypes.ResponseCheckTx{Code: abcitypes.CodeTypeOK}, nil
}
