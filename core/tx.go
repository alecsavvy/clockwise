/* utilities specific to the core, not intended for external use */
package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cometbft/cometbft/types"
)

func (c *Core) Send(tx interface{}) (*ManageEntity, error) {
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
		var res ManageEntity
		err := c.fromTxBytes(etx.Tx, &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
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

func (c *Core) fromTxBytes(jsonBytes []byte, result interface{}) error {
	err := json.Unmarshal(jsonBytes, result)
	if err != nil {
		return err
	}

	return nil
}
