package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/alecsavvy/clockwise/api"
	"github.com/labstack/echo/v4"
)

type HealthResponses struct {
	Responses []*api.HealthCheckResponse
}

func getHealthStats(c echo.Context) error {
	var healthResponses []*api.HealthCheckResponse
	for _, endpoint := range discprovUrls {
		res, err := http.Get(fmt.Sprintf("%s/health_check", endpoint))
		if err != nil {
			log.Printf("error getting health %s %v", endpoint, err)
			continue
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error reading response from %s: %v", endpoint, err)
			continue
		}

		var healthResp api.HealthCheckResponse
		if err := json.Unmarshal(body, &healthResp); err != nil {
			log.Printf("Error unmarshalling response from %s: %v", endpoint, err)
			continue
		}

		healthResponses = append(healthResponses, &healthResp)
	}

	h := &HealthResponses{
		Responses: healthResponses,
	}
	return health_templ.ExecuteTemplate(c.Response().Writer, "health", h)
}
