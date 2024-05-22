package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/core/interface/commands"
	"github.com/alecsavvy/clockwise/core/interface/entities"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func (c *Core) CreateFollow(cmd *commands.CreateFollowCommand) (*entities.FollowEntity, error) {
	ctx := context.Background()
	queries := c.db

	_, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}

	follow, err := queries.GetFollowByID(ctx, cmd.Data.ID)
	if err != nil {
		return nil, err
	}

	followEntity := c.followModelsToEntities([]db.Follow{follow})[0]
	return followEntity, nil
}

func (c *Core) finalizeCreateFollow(qtx *db.Queries, createdAt int64, b []byte) (*abcitypes.ExecTxResult, error) {
	ctx := context.Background()

	var cmd commands.CreateFollowCommand
	err := c.fromTxBytes(b, &cmd)
	if err != nil {
		return nil, utils.AppError("not a create follo command in create follow handler", err)
	}

	follo := cmd.Data

	err = qtx.CreateFollow(ctx, db.CreateFollowParams{
		ID:          follo.ID,
		FollowerID:  follo.FollowerID,
		FollowingID: follo.FollowingID,
		CreatedAt:   createdAt,
	})
	return &abcitypes.ExecTxResult{
		Code: 0,
	}, nil
}

func (c *Core) GetUserFollowers(userId string) ([]*entities.FollowEntity, error) {
	ctx := context.Background()
	db := c.db

	follows, err := db.GetFollowers(ctx, userId)
	if err != nil {
		return nil, err
	}

	return c.followModelsToEntities(follows), nil
}

func (c *Core) GetUserFollowing(userId string) ([]*entities.FollowEntity, error) {
	ctx := context.Background()
	db := c.db

	follows, err := db.GetFollowing(ctx, userId)
	if err != nil {
		return nil, err
	}

	return c.followModelsToEntities(follows), nil
}
