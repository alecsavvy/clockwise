package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/pubsub"
	ctypes "github.com/cometbft/cometbft/types"
)

type UserPubsub = pubsub.Pubsub[*entities.UserEntity]
type TrackPubsub = pubsub.Pubsub[*entities.TrackEntity]
type FollowPubsub = pubsub.Pubsub[*entities.FollowEntity]
type RepostPubsub = pubsub.Pubsub[*entities.RepostEntity]

type Pubsub struct {
	UserPubsub   *UserPubsub
	TrackPubsub  *TrackPubsub
	FollowPubsub *FollowPubsub
	RepostPubsub *RepostPubsub
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		UserPubsub:   pubsub.NewPubsub[*entities.UserEntity](),
		TrackPubsub:  pubsub.NewPubsub[*entities.TrackEntity](),
		FollowPubsub: pubsub.NewPubsub[*entities.FollowEntity](),
		RepostPubsub: pubsub.NewPubsub[*entities.RepostEntity](),
	}
}

func (c *Core) RunPubsub() error {
	rpc := c.rpc

	err := rpc.Start()
	if err != nil {
		return err
	}
	defer rpc.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventCh, err := rpc.Subscribe(ctx, "block-subscriber", "tm.event = 'NewBlock'")

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
	ctx := context.Background()
	for _, tx := range txs {
		var cmd commands.Command[any]
		c.fromTxBytes(tx, &cmd)
		operation := cmd.Operation

		switch operation {
		case commands.Operation{Action: commands.CREATE, Entity: commands.USER}:
			var e commands.CreateUserCommand
			c.fromTxBytes(tx, &e)
			user, _ := c.db.GetUserByID(ctx, e.Data.ID)
			userEntity := c.userModelsToEntities([]db.User{user})[0]
			c.pubsub.UserPubsub.Publish(userEntity)
		case commands.Operation{Action: commands.CREATE, Entity: commands.TRACK}:
			var e commands.CreateTrackCommand
			c.fromTxBytes(tx, &e)
			track, _ := c.db.GetTrackByID(ctx, e.Data.ID)
			trackEntity := c.trackModelsToEntities([]db.Track{track})[0]
			c.pubsub.TrackPubsub.Publish(trackEntity)
		case commands.Operation{Action: commands.CREATE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.CREATE, Entity: commands.REPOST}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.REPOST}:
		default:
		}
	}
}
