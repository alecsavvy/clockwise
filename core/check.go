/*
validate_tx.go

This file contains the logic for validating transactions at various states in the chains consensus.
*/

package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Performs validation on a proposed transaction, should be very performant as this check
// gets called a lot (per the cometbft docs)
func (c *Core) CheckTx(ctx context.Context, req *abcitypes.RequestCheckTx) (*abcitypes.ResponseCheckTx, error) {
	return &abcitypes.ResponseCheckTx{Code: abcitypes.CodeTypeOK}, nil
}

func successCheck() (*abcitypes.ResponseCheckTx, error) {
	return &abcitypes.ResponseCheckTx{Code: abcitypes.CodeTypeOK}, nil
}

func errorCheck() (*abcitypes.ResponseCheckTx, error) {
	return &abcitypes.ResponseCheckTx{Code: 1}, nil
}
