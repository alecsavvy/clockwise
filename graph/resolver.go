package graph

import (
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	core    *core.Core
	logger  *utils.Logger
	queries *db.Queries
}

func NewResolver(logger *utils.Logger, core *core.Core, pool *pgxpool.Pool) *Resolver {
	return &Resolver{
		core:    core,
		logger:  logger,
		queries: db.New(pool),
	}
}
