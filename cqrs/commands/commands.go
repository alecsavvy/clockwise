package commands

import "github.com/google/uuid"

type CommandType = string

var (
	CreateUser   CommandType = "CreateUser"
	CreateTrack  CommandType = "CreateTrack"
	CreateFollow CommandType = "CreateFollow"
	CreateRepost CommandType = "CreateRepost"
)

type Command struct {
	CommandType
}

type CreateUserCommand struct {
	Command
	ID      string
	Handle  string
	Address string
	Bio     string
}

type CreateTrackCommand struct {
	Command
	Title       string
	StreamURL   string
	Description string
	UserID      uuid.UUID
}

type CreateFollowCommand struct {
	Command
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
}

type CreateRepostCommand struct {
	Command
	ReposterID uuid.UUID
	TrackID    uuid.UUID
}
