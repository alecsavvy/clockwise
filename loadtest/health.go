package main

import (
	"context"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	NodeID string
	IsHealthy string
	UnconfirmedTxs uint
	LatestBlockHeight uint
	CatchingUp bool
	Peers uint 
}

type HealthResponses struct {
	Responses []*HealthResponse
}

func getHealthStats(c echo.Context) error {
	ctx := context.Background()

	var healthResponses []*HealthResponse
	for _, endpoint := range discprovUrls {
		rpcEndpoint := string(append([]rune(endpoint)[:len([]rune(endpoint))-1], '7'))

		// create rpc connection
		rpc, err := http.New(rpcEndpoint)
		if err != nil {
			return err
		}

		status, err := rpc.Status(ctx)
		if err != nil {
			return err
		}

		unconfirmed, err := rpc.UnconfirmedTxs(ctx, nil)
		if err != nil {
			return err
		}
	
		netinfo, err := rpc.NetInfo(ctx)
		if err != nil {
			return err
		}

		healthResp := &HealthResponse{
			NodeID: string(status.NodeInfo.ID()),
			LatestBlockHeight: uint(status.SyncInfo.LatestBlockHeight),
			CatchingUp: status.SyncInfo.CatchingUp,
			Peers: uint(netinfo.NPeers),
			UnconfirmedTxs: uint(unconfirmed.Total),
			IsHealthy: "true",
		}

		healthResponses = append(healthResponses, healthResp)
	}

	h := &HealthResponses{
		Responses: healthResponses,
	}
	return health_templ.ExecuteTemplate(c.Response().Writer, "health", h)
}
