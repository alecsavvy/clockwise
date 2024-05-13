package adapters

import (
	"context"

	chainclient "github.com/alecsavvy/clockwise/core/chain_client"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/cqrs/events"
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/utils"
)

type UserRepository struct {
	logger *utils.Logger
	cc     *chainclient.ChainClient
	db     *db.Queries
}

func NewUserRepo(logger *utils.Logger, cc *chainclient.ChainClient, db *db.Queries) *UserRepository {
	return &UserRepository{
		logger: logger,
		cc:     cc,
		db:     db,
	}
}

func (ur *UserRepository) CreateUser(cmd *commands.CreateUserCommand) (*events.UserCreatedEvent, error) {
	ctx := context.Background()
	cc := ur.cc
	db := ur.db

	// submit command to chain
	res, err := cc.Send(cmd)
	if err != nil {
		return nil, err
	}

	// construct event
	var event events.UserCreatedEvent
	event.BlockHeight = uint64(res.Height)
	event.TransactionHash = string(res.Hash)

	user, err := db.GetUserByHandle(ctx, cmd.Handle)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.UserEntity{
		ID:      user.ID,
		Handle:  user.Handle,
		Bio:     user.Bio,
		Address: user.Address,
	}

	event.User = *userEntity

	return &event, nil
}

func (ur *UserRepository) GetUsers() ([]*entities.UserEntity, error) {
	ctx := context.Background()

	users, err := ur.db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	userEntities := utils.Map(users, func(user db.User) *entities.UserEntity {
		return &entities.UserEntity{
			ID:      user.ID,
			Handle:  user.Handle,
			Bio:     user.Bio,
			Address: user.Address,
		}
	})

	return userEntities, nil
}

var _ services.UserService = (*UserRepository)(nil)
