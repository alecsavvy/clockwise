package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
)

// returns in current postgres tx for this block
func (c *Core) getDb() *db.Queries {
	return c.queries.WithTx(c.currentTx)
}

func (c *Core) startInProgressTx(ctx context.Context) error {
	dbTx, err := c.pool.Begin(ctx)
	if err != nil {
		return err
	}

	c.currentTx = dbTx
	return nil
}

// commits the current tx that's finished indexing
func (c *Core) commitInProgressTx(ctx context.Context) error {
	if c.currentTx != nil {
		err := c.currentTx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
