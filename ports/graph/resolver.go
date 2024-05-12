//go:generate go run github.com/99designs/gqlgen generate

package graph

import (
	"github.com/alecsavvy/clockwise/ports/graph/model"
	"github.com/alecsavvy/clockwise/utils"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
)

type UserDB = []*model.User
type TrackDB = []*model.Track

type Resolver struct {
	users       UserDB
	tracks      TrackDB
	chainClient *rpchttp.HTTP
	logger      *utils.Logger
}

func NewResolver(logger *utils.Logger, chainClient *rpchttp.HTTP) *Resolver {
	return &Resolver{
		chainClient: chainClient,
		logger:      logger,
	}
}
