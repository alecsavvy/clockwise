package api

import (
	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/utils"
)

type Api struct {
	core   *core.Core
	logger *utils.Logger
}

func NewApi(logger *utils.Logger, core *core.Core) *Api {
	return &Api{
		logger: logger,
		core:   core,
	}
}
