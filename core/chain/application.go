package chain

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Clockwise implementation of the abci interface for cometbft
// everything related to consensus MUST go through here
// https://docs.cometbft.com/v0.38/spec/abci/abci++_methods
type Application struct {
	logger    *utils.Logger
	queries   *db.Queries
	pool      *pgxpool.Pool
	currentTx pgx.Tx
}

// compile time check for abci compatibility
var _ abcitypes.Application = (*Application)(nil)

func NewApplication(logger *utils.Logger, queries *db.Queries, pool *pgxpool.Pool) *Application {
	return &Application{
		logger:  logger,
		queries: queries,
		pool:    pool,
	}
}
