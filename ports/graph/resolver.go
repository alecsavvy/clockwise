//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/core/client"
	"github.com/alecsavvy/clockwise/utils"
)

type Resolver struct {
	core   *client.Core
	logger *utils.Logger
}

func NewResolver(logger *utils.Logger, core *client.Core) *Resolver {
	return &Resolver{
		core:   core,
		logger: logger,
	}
}
