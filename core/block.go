package core

import (
	"context"

	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func CheckTx(c *Core, ctx context.Context, req *abcitypes.RequestCheckTx) error {
	return nil
}
