package api

import (
	"fmt"

	ctypes "github.com/cometbft/cometbft/types"
	"github.com/labstack/echo/v4"
)

func (api *Api) SSEHandler(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")

	// Ensure that the server sends a response immediately
	c.Response().Flush()

	blockChan := make(chan *ctypes.Block)
	go func() {
		defer close(blockChan)
		blocks := api.core.Pubsub().NewBlockPubsub.Subscribe()
		for {
			select {
			case newBlock, ok := <-blocks:
				if !ok {
					return
				}
				blockChan <- newBlock
			case <-c.Request().Context().Done():
				return
			}
		}
	}()

	quit := make(chan struct{})
	go func() {
		<-c.Request().Context().Done()
		close(quit)
	}()

	for {
		select {
		case <-quit:
			return nil
		case block := <-blockChan:
			// Write the SSE data
			fmt.Fprintf(c.Response(), "data: %v\n\n", block.Height)
			c.Response().Flush()
		}
	}
}
