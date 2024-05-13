//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/ports/graph/model"
	"github.com/alecsavvy/clockwise/utils"
)

type TrackDB = []*model.Track

type Resolver struct {
	tracks      TrackDB
	userService services.UserService
	logger      *utils.Logger
}

func NewResolver(logger *utils.Logger, userService services.UserService) *Resolver {
	return &Resolver{
		userService: userService,
		logger:      logger,
	}
}
