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
	/** writes */
	CreateUser(*commands.CreateUserCommand) (*events.UserCreatedEvent, error)
	FollowUser(*commands.CreateFollowCommand) (*events.FollowCreatedEvent, error)
	UnfollowUser(followId string) (string, error)

	/** reads */
	GetUsers() ([]*entities.UserEntity, error)
	GetUserFollowing(userId string) ([]*entities.FollowEntity, error)
	GetUserFollowers(userId string) ([]*entities.FollowEntity, error)
	GetUserReposts(userId string) ([]*entities.RepostEntity, error)
}

type TrackService interface {
	/** writes */
	CreateTrack(*commands.CreateTrackCommand) (*events.TrackCreatedEvent, error)
	RepostTrack(*commands.CreateRepostCommand) (*events.RepostCreatedEvent, error)

	/** reads */
	GetTracks() ([]*entities.TrackEntity, error)
	GetTrackReposts(trackId string) ([]*entities.RepostEntity, error)
}
