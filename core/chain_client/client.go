package chainclient

import (
	"context"
	"time"

	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/utils"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

type ChainClient struct {
	rpc    *rpchttp.HTTP
	logger *utils.Logger
}

func New(logger *utils.Logger, rpc *rpchttp.HTTP) *ChainClient {
	return &ChainClient{
		logger: logger,
		rpc:    rpc,
	}
}

func (cc *ChainClient) Send(tx interface{}) (*ctypes.ResultTx, error) {
	ctx := context.Background()
	rpc := cc.rpc

	txBytes, err := chainutils.ToTxBytes(tx)
	if err != nil {
		return nil, err
	}

	result, err := rpc.BroadcastTxSync(ctx, txBytes)
	if err != nil {
		return nil, utils.AppError("failure to broadcast tx", err)
	}

	for {
		tx, err := rpc.Tx(ctx, result.Hash, true)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		return tx, nil
	}
}

func (cc *ChainClient) GetRpc() *rpchttp.HTTP {
	return cc.rpc
}
