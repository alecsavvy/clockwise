package core

import (
	"context"

	"github.com/alecsavvy/clockwise/pubsub"
	ctypes "github.com/cometbft/cometbft/types"
)

type EntityManagerPubsub = pubsub.Pubsub[*ManageEntity]
type NewBlockPubsub = pubsub.Pubsub[*ctypes.Block]

type Pubsub struct {
	EntityManagerPubsub *EntityManagerPubsub
	NewBlockPubsub      *NewBlockPubsub
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		EntityManagerPubsub: pubsub.NewPubsub[*ManageEntity](),
		NewBlockPubsub:      pubsub.NewPubsub[*ctypes.Block](),
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

			// publish full blocks to listeneres
			c.pubsub.NewBlockPubsub.Publish(block.Block)

			txs := block.Block.Txs
			for _, tx := range txs {
				// publish manage entities to listeners
				var me ManageEntity
				c.fromTxBytes(tx, &me)
				c.pubsub.EntityManagerPubsub.Publish(&me)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
