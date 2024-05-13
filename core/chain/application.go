package chain

import (
	"github.com/alecsavvy/clockwise/core/db"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Clockwise implementation of the abci interface for cometbft
// everything related to consensus MUST go through here
// https://docs.cometbft.com/v0.38/spec/abci/abci++_methods
type Application struct {
	queries   *db.Queries
	pool      *pgxpool.Pool
	currentTx pgx.Tx
}

// compile time check for abci compatibility
var _ abcitypes.Application = (*Application)(nil)

func NewApplication() *Application {
	return &Application{}
}
