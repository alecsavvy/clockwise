package api

import (
	"context"
	"errors"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/labstack/echo/v4"
)

type ErrorBucket struct {
	bucket []error
}

func (eb *ErrorBucket) addError(msg string, err error) {
	eb.bucket = append(eb.bucket, errors.Join(errors.New(msg), err))
}

type HealthCheckResponse struct {
	IsHealthy bool                     `json:"is_healthy"`
	NetInfo   *coretypes.ResultNetInfo `json:"net_info"`
	Status    *coretypes.ResultStatus  `json:"core_status"`
	Errors    ErrorBucket              `json:"errors"`
}

func (api *Api) HealthCheck(c echo.Context) error {
	ctx := context.Background()
	logger := api.logger
	core := api.core
	rpc := core.Rpc()

	var res HealthCheckResponse
	var errorBucket ErrorBucket

	netinfo, err := rpc.NetInfo(ctx)
	if err != nil {
		errorBucket.addError("unhealthy netinfo", err)
	}

	status, err := rpc.Status(ctx)
	if err != nil {
		errorBucket.addError("unhealthy core status", err)
	}

	res.NetInfo = netinfo
	res.Status = status
	res.Errors = errorBucket
	res.IsHealthy = len(errorBucket.bucket) == 0

	if !res.IsHealthy {
		for _, err := range res.Errors.bucket {
			logger.Error("health check errors", "err", err)
		}
	}

	return c.JSONPretty(200, res, "  ")
}
