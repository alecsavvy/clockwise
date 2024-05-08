package entities

import "github.com/google/uuid"

type TrackEntity struct {
	ID          uuid.UUID
	Title       string
	Description string
	StreamURL   string
	UserID      uuid.UUID
}
