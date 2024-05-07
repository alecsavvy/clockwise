//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/alecsavvy/clockwise/graph/model"

type UserDB = []*model.User
type TrackDB = []*model.Track

type Resolver struct {
	users  UserDB
	tracks TrackDB
}

func NewResolver() *Resolver {
	return &Resolver{}
}
