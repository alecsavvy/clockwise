/*
services.go

Put all services interfaces in one file unless they need to be broken out.
*/
package services

import (
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/cqrs/events"
)

type UserService interface {
	CreateUser(*commands.CreateUserCommand) (*events.UserCreatedEvent, error)
	GetUserByHandle(handle string) (*entities.UserEntity, error)
	GetUsers() ([]*entities.UserEntity, error)
}

type TrackService interface {
	CreateTrack(*commands.CreateTrackCommand) (*events.TrackCreatedEvent, error)
	GetTrackByTitle(title string) (*entities.TrackEntity, error)
	GetTracks() ([]*entities.TrackEntity, error)
}
