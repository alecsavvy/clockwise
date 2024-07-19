package main

import (
	"fmt"
	"strconv"

	"github.com/alecsavvy/clockwise/core"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

type TxView struct {
	MessageType string
	Signature string
	TxHash string
}

type Block struct {
	Rpc            string
	BlockNumber    int
	BlockHash      string
	TotalTxs       int
	Txs []*TxView
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
	var txs []*TxView
	for _, tx := range block.Block.Txs {
		var ev gen.Envelope
		err := proto.Unmarshal(tx, &ev)
		if err != nil {
			logger.Error(err, "uh oh")
			return err
		}

		txhash, err := core.ToTxHash(&ev)
		if err != nil {
			logger.Error(err, "uh oh")
			return err
		}

		view := &TxView{
			MessageType: ev.Headers.GetMessageType().String(),
			TxHash: txhash,
			Signature: ev.Headers.GetSignature(),
		}

		txs = append(txs, view)
	}

	// render page with template
	b := &Block{
		Rpc:            endpoint,
		BlockNumber:    blockNumber,
		BlockHash:      blockHash,
		TotalTxs:       len(txs),
		Txs: txs,
	}
	return block_templ.ExecuteTemplate(c.Response().Writer, "block", b)
}
