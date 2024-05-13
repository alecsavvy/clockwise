package adapters

import (
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/cqrs/events"
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/utils"
)

type UserRepository struct{}

func (ur *UserRepository) CreateUser(cmd *commands.CreateUserCommand) (*events.UserCreatedEvent, error) {
	return nil, utils.AppError("not implemented", nil)
}

func (ur *UserRepository) GetUserByHandle(handle string) (*entities.UserEntity, error) {
	return nil, utils.AppError("not implemented", nil)
}

func (ur *UserRepository) GetUsers() ([]*entities.UserEntity, error) {
	return nil, utils.AppError("not implemented", nil)
}

func NewUserRepo() *UserRepository {
	return &UserRepository{}
}

var _ services.UserService = (*UserRepository)(nil)
