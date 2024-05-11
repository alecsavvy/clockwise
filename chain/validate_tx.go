/*
validate_tx.go

This file contains the logic for validating transactions at various states in the chains consensus.
*/

package chain

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (a *Application) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	_ = req.GetTx()

	return &abcitypes.ResponseCheckTx{}, nil
}
