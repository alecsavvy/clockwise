//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/utils"
)

type Resolver struct {
	core   *core.Core
	logger *utils.Logger
}

func NewResolver(logger *utils.Logger, core *core.Core) *Resolver {
	return &Resolver{
		core:   core,
		logger: logger,
	}
}
