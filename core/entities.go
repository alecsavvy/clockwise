package core

import "github.com/google/uuid"

type UserEntity struct {
	ID     *uuid.UUID
	Handle string
	Bio    string
}

type TrackEntity struct {
	ID          *uuid.UUID
	Title       string
	StreamURL   string
	Description string
	UserID      *uuid.UUID
}
