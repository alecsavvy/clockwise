/* utilities specific to the core, not intended for external use */
package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/types"
)

func (c *Core) Send(tx interface{}) (*abcitypes.TxResult, error) {
	ctx := context.Background()
	rpc := c.rpc

	txBytes, err := c.toTxBytes(tx)
	if err != nil {
		return nil, err
	}

	result, err := rpc.BroadcastTxSync(ctx, txBytes)
	if err != nil {
		return nil, err
	}

	txChan, err := rpc.Subscribe(ctx, "tx-subscriber", fmt.Sprintf("tm.event = 'Tx' AND tx.hash = '%X'", result.Hash))
	if err != nil {
		return nil, err
	}

	select {
	case tx := <-txChan:
		etx := tx.Data.(types.EventDataTx)
		return &etx.TxResult, nil
	case <-time.After(30 * time.Second):
		return nil, errors.New("tx waiting timeout")
	}
}

func (c *Core) toTxBytes(tx interface{}) ([]byte, error) {
	txBytes, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

func (c *Core) fromTxBytes(jsonBytes []byte) (interface{}, error) {
	var result interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
