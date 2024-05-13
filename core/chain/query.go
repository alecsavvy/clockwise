/*
query.go

Contains logic to directly query the chain. This may be irrelevant with gql having a direct handle on the db.
*/
package chain

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (a *Application) Query(_ context.Context, req *abcitypes.RequestQuery) (*abcitypes.ResponseQuery, error) {
	resp := abcitypes.ResponseQuery{Key: req.Data}
	return &resp, nil
}
