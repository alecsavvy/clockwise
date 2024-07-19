package main

import (
	"fmt"
	"strconv"

	"github.com/alecsavvy/clockwise/core"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/labstack/echo/v4"
)

type Block struct {
	Rpc            string
	BlockNumber    int
	BlockHash      string
	TotalTxs       int
	ManageEntities []*core.ManageEntity
}

func getBlock(c echo.Context) error {
	logger := c.Logger()
	ctx := c.Request().Context()
	// parse out form
	option := c.FormValue("option")
	blockNumber, err := strconv.Atoi(c.FormValue("Block Number"))
	blockHeight := int64(blockNumber)
	if err != nil {
		return err
	}
	endpoint := fmt.Sprintf("http://node-%s:26657", option)

	// create rpc connection
	client, err := http.New(endpoint)
	if err != nil {
		logger.Error(err, "could not create rpc")
		return err
	}

	// get block
	block, err := client.Block(ctx, &blockHeight)
	if err != nil {
		logger.Error(err, "could not get block, doesn't exist or was pruned")
		return err
	}

	blockHash := block.Block.Hash().String()

	// parse out em txs
	var manageEntities []*core.ManageEntity
	// for _, tx := range block.Block.Txs {
	// 	var em core.ManageEntity
	// 	err := json.Unmarshal(tx, &em)
	// 	if err != nil {
	// 		logger.Error(err, "uh oh")
	// 		return err
	// 	}
	// 	manageEntities = append(manageEntities, &em)
	// }

	// render page with template
	b := &Block{
		Rpc:            endpoint,
		BlockNumber:    blockNumber,
		BlockHash:      blockHash,
		TotalTxs:       len(manageEntities),
		ManageEntities: manageEntities,
	}
	return block_templ.ExecuteTemplate(c.Response().Writer, "block", b)
}
