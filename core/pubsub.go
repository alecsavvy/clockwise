package core

import (
	"context"
	"fmt"

	"github.com/alecsavvy/clockwise/protocol"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/alecsavvy/clockwise/pubsub"
	ctypes "github.com/cometbft/cometbft/types"
	"google.golang.org/protobuf/proto"
)

type EntityManagerPubsub = pubsub.Pubsub[*ManageEntity]
type NewBlockPubsub = pubsub.Pubsub[*ctypes.Block]

type CreateUserPubsub = pubsub.Pubsub[*gen.CreateUser]
type CreateTrackPubsub = pubsub.Pubsub[*gen.CreateTrack]
type FollowUserPubsub = pubsub.Pubsub[*gen.FollowUser]
type RepostTrackPubsub = pubsub.Pubsub[*gen.RepostTrack]

type Pubsub struct {
	CreateUserPubsub  *CreateUserPubsub
	CreateTrackPubsub *CreateTrackPubsub
	FollowUserPubsub  *FollowUserPubsub
	RepostTrackPubsub *RepostTrackPubsub
	NewBlockPubsub    *NewBlockPubsub

	pubsubRoutes *protocol.MessageRouterMap
}

func NewPubsub() *Pubsub {
	ps := &Pubsub{
		CreateUserPubsub:  pubsub.NewPubsub[*gen.CreateUser](),
		CreateTrackPubsub: pubsub.NewPubsub[*gen.CreateTrack](),
		FollowUserPubsub:  pubsub.NewPubsub[*gen.FollowUser](),
		RepostTrackPubsub: pubsub.NewPubsub[*gen.RepostTrack](),
		NewBlockPubsub:    pubsub.NewPubsub[*ctypes.Block](),
	}

	pubsubRoutes := make(protocol.MessageRouterMap, 0)
	pubsubRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_USER] = ps.createUserPublish
	pubsubRoutes[gen.MessageType_MESSAGE_TYPE_CREATE_TRACK] = ps.createTrackPublish
	pubsubRoutes[gen.MessageType_MESSAGE_TYPE_REPOST_TRACK] = ps.repostTrackPublish
	pubsubRoutes[gen.MessageType_MESSAGE_TYPE_FOLLOW_USER] = ps.followUserPublish

	ps.pubsubRoutes = &pubsubRoutes

	return ps
}

func (ps *Pubsub) createUserPublish(ctx context.Context, msg proto.Message) error {
	ps.CreateUserPubsub.Publish(ctx, msg.(*gen.CreateUser))
	return nil
}

func (ps *Pubsub) createTrackPublish(ctx context.Context, msg proto.Message) error {
	ps.CreateTrackPubsub.Publish(ctx, msg.(*gen.CreateTrack))
	return nil
}

func (ps *Pubsub) followUserPublish(ctx context.Context, msg proto.Message) error {
	ps.FollowUserPubsub.Publish(ctx, msg.(*gen.FollowUser))
	return nil
}

func (ps *Pubsub) repostTrackPublish(ctx context.Context, msg proto.Message) error {
	ps.RepostTrackPubsub.Publish(ctx, msg.(*gen.RepostTrack))
	return nil
}

func (c *Core) RunPubsub() error {
	rpc := c.rpc

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh, err := rpc.Subscribe(ctx, "block-subscriber", "tm.event = 'NewBlock'")
	if err != nil {
		return err
	}

	defer func() {
		if err := rpc.Unsubscribe(ctx, "block-subscriber", "tm.event = 'NewBlock'"); err != nil {
			// Handle the unsubscribe error if necessary
			fmt.Println("Failed to unsubscribe:", err)
		}
	}()

	for {
		select {
		case event := <-eventCh:
			block, ok := event.Data.(ctypes.EventDataNewBlock)
			if !ok {
				continue
			}

			// publish full blocks to listeners
			c.pubsub.NewBlockPubsub.Publish(ctx, block.Block)

			txs := block.Block.Txs
			for _, tx := range txs {
				// broadcast specific messages
				go protocol.MessageRouter(ctx, *c.pubsub.pubsubRoutes, tx)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
