package client

import (
	"context"

	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
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
