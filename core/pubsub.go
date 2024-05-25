package core

import (
	"context"

	"github.com/alecsavvy/clockwise/pubsub"
	ctypes "github.com/cometbft/cometbft/types"
)

type EntityManagerPubsub = pubsub.Pubsub[*ManageEntity]

type Pubsub struct {
	EntityManagerPubsub *EntityManagerPubsub
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		EntityManagerPubsub: pubsub.NewPubsub[*ManageEntity](),
	}
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
			for _, tx := range txs {
				var me ManageEntity
				c.fromTxBytes(tx, &me)
				c.pubsub.EntityManagerPubsub.Publish(&me)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
