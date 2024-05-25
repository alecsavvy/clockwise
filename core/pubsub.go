package core

import (
	"context"

	ctypes "github.com/cometbft/cometbft/types"
)

type Pubsub struct{}

func NewPubsub() *Pubsub {
	return &Pubsub{}
}

func (c *Core) RunPubsub() error {
	rpc := c.rpc

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh, err := rpc.Subscribe(ctx, "block-subscriber", "tm.event = 'NewBlock'")
	if err != nil {
		return err
	}

	for {
		select {
		case event := <-eventCh:
			block, ok := event.Data.(ctypes.EventDataNewBlock)
			if !ok {
				continue
			}

			txs := block.Block.Txs
			c.broadcastTxs(txs)
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *Core) broadcastTxs(txs ctypes.Txs) {
	// context.Background()
	// for _, _ := range txs {
	// 	// broadcast txs
	// }
}
