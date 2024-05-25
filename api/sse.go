package api

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func (api *Api) SSEHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")

	// Ensure that the server sends a response immediately
	c.Response().Flush()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	quit := make(chan struct{})
	go func() {
		<-c.Request().Context().Done()
		close(quit)
	}()

	for {
		select {
		case <-quit:
			return nil
		case t := <-ticker.C:
			// Write the SSE data
			fmt.Fprintf(c.Response(), "data: %v\n\n", t.Format(time.RFC1123))
			c.Response().Flush()
		}
	}
}
