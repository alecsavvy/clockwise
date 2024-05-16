package commands

import "github.com/google/uuid"

type CommandType = string

// command actions
var (
	CREATE = "Create"
	UPDATE = "Update"
	DELETE = "Delete"
)

// command entities
var (
	USER   = "User"
	TRACK  = "Track"
	FOLLOW = "Follow"
	REPOST = "Repost"
)

type Operation struct {
	Entity string
	Action string
}

type Command[T any] struct {
	Operation
	Data T
}

func NewCommand[T any](entity, action string, data T) *Command[T] {
	return &Command[T]{
		Operation: Operation{
			Entity: entity,
			Action: action,
		},
		Data: data,
	}
}

type CreateUserCommand = Command[CreateUser]
type CreateTrackCommand = Command[CreateTrack]
type CreateFollowCommand = Command[CreateFollow]
type CreateRepostCommand = Command[CreateRepost]

type CreateUser struct {
	ID      string
	Handle  string
	Address string
	Bio     string
}

type CreateTrack struct {
	Title       string
	StreamURL   string
	Description string
	UserID      uuid.UUID
}

type CreateFollow struct {
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
}

type CreateRepost struct {
	ReposterID uuid.UUID
	TrackID    uuid.UUID
}
