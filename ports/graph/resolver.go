//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/ports/graph/model"
	"github.com/alecsavvy/clockwise/utils"
)

type TrackDB = []*model.Track

type Resolver struct {
	trackService services.TrackService
	userService  services.UserService
	logger       *utils.Logger
}

func NewResolver(logger *utils.Logger, userService services.UserService, trackService services.TrackService) *Resolver {
	return &Resolver{
		trackService: trackService,
		userService:  userService,
		logger:       logger,
	}
}
