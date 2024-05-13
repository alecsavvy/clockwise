package entities

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	ID      uuid.UUID
	Handle  string
	Address string
	Bio     string
}

type TrackEntity struct {
	ID          uuid.UUID
	Title       string
	StreamURL   string
	Description string
	UserID      uuid.UUID
}

type FollowEntity struct {
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
}

type RepostEntity struct {
	ReposterID uuid.UUID
	TrackID    uuid.UUID
}
