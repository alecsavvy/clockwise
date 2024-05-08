package entities

import "github.com/google/uuid"

type UserEntity struct {
	ID     uuid.UUID
	Handle string
	Bio    string
}
