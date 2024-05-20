/*
query.go

Contains logic to directly query the chain. This may be irrelevant with gql having a direct handle on the db.
*/
package core

import (
	"context"
	"errors"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

// Reads from the chain state, may be irrelevant with the higher level GQL interface
func (c *Core) Query(_ context.Context, req *abcitypes.RequestQuery) (*abcitypes.ResponseQuery, error) {
	return nil, errors.New("internal query not supported, please use higher level ports")
}
