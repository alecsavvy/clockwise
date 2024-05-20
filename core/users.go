package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/core/interface/commands"
	"github.com/alecsavvy/clockwise/core/interface/entities"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func (c *Core) CreateUser(cmd *commands.CreateUserCommand) (*entities.UserEntity, error) {
	ctx := context.Background()
	db := c.db

	// submit command to chain
	_, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}

	user, err := db.GetUserByHandle(ctx, cmd.Data.Handle)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.UserEntity{
		ID:      user.ID,
		Handle:  user.Handle,
		Bio:     user.Bio,
		Address: user.Address,
	}

	return userEntity, nil
}

func (c *Core) handleCreateUser(qtx *db.Queries, createdAt int64, b []byte) (*abcitypes.ExecTxResult, error) {
	ctx := context.Background()

	var cmd commands.CreateUserCommand
	err := c.fromTxBytes(b, &cmd)
	if err != nil {
		return nil, utils.AppError("not a create user command in create user handler", err)
	}

	user := cmd.Data

	err = qtx.CreateUser(ctx, db.CreateUserParams{
		ID:        user.ID,
		Handle:    user.Handle,
		Bio:       user.Bio,
		Address:   user.Address,
		CreatedAt: createdAt,
	})

	if err != nil {
		return nil, utils.AppError("failure to insert user", err)
	}

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

func (c *Core) GetUserReposts(userId string) ([]*entities.RepostEntity, error) {
	ctx := context.Background()
	db := c.db

	reposts, err := db.GetUserReposts(ctx, userId)
	if err != nil {
		return nil, err
	}

	return c.repostModelsToEntities(reposts), nil
}

func (c *Core) GetUsers() ([]*entities.UserEntity, error) {
	ctx := context.Background()

	users, err := c.db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return c.userModelsToEntities(users), nil
}
