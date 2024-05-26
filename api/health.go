package api

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
)

type ErrorBucket struct {
	bucket []error
}

func (eb *ErrorBucket) addError(msg string, err error) {
	eb.bucket = append(eb.bucket, errors.Join(errors.New(msg), err))
}

type HealthCheckResponse struct {
	IsHealthy         bool        `json:"is_healthy"`
	NodeID            string      `json:"node_id"`
	LatestBlockHeight uint        `json:"latest_block_height"`
	CatchingUp        bool        `json:"catching_up"`
	Peers             uint        `json:"peers"`
	Errors            ErrorBucket `json:"errors"`
}

func (api *Api) HealthCheck(c echo.Context) error {
	ctx := context.Background()
	logger := api.logger
	core := api.core
	rpc := core.Rpc()

	var res HealthCheckResponse
	var errorBucket ErrorBucket

	status, err := rpc.Status(ctx)
	if err != nil {
		errorBucket.addError("unhealthy core status", err)
	}

	netinfo, err := rpc.NetInfo(ctx)
	if err != nil {
		errorBucket.addError("unhealthy net info", err)
	}

	res.NodeID = string(status.NodeInfo.ID())
	res.LatestBlockHeight = uint(status.SyncInfo.LatestBlockHeight)
	res.CatchingUp = status.SyncInfo.CatchingUp
	res.Peers = uint(netinfo.NPeers)
	res.Errors = errorBucket
	res.IsHealthy = len(errorBucket.bucket) == 0

	if !res.IsHealthy {
		for _, err := range res.Errors.bucket {
			logger.Error("health check errors", "err", err)
		}
	}

	return c.JSONPretty(200, res, "  ")
}
