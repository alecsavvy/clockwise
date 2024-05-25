//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

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
