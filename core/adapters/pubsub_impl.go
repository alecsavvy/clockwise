package adapters

import (
	"context"

	chainclient "github.com/alecsavvy/clockwise/core/chain_client"
	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/pubsub"
	ctypes "github.com/cometbft/cometbft/types"
)

type PubsubAdapter struct {
	cc *chainclient.ChainClient
	db *db.Queries

	UserPubsub   *services.UserPubsub
	TrackPubsub  *services.TrackPubsub
	FollowPubsub *services.FollowPubsub
	RepostPubsub *services.RepostPubsub
}

func NewPubsubAdapter(cc *chainclient.ChainClient, db *db.Queries) *PubsubAdapter {
	return &PubsubAdapter{
		cc:           cc,
		db:           db,
		UserPubsub:   pubsub.NewPubsub[*entities.UserEntity](),
		TrackPubsub:  pubsub.NewPubsub[*entities.TrackEntity](),
		FollowPubsub: pubsub.NewPubsub[*entities.FollowEntity](),
		RepostPubsub: pubsub.NewPubsub[*entities.RepostEntity](),
	}
}

func (ps *PubsubAdapter) Run() error {
	cc := ps.cc
	rpc := cc.GetRpc()

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
			ps.broadcastTxs(txs)
		case <-ctx.Done():
			return nil
		}
	}
}

func (ps *PubsubAdapter) broadcastTxs(txs ctypes.Txs) {
	ctx := context.Background()
	for _, tx := range txs {
		cmd, _ := chainutils.FromTxBytes[commands.Command[any]](tx)
		operation := cmd.Operation

		switch operation {
		case commands.Operation{Action: commands.CREATE, Entity: commands.USER}:
			e, _ := chainutils.FromTxBytes[commands.CreateUserCommand](tx)
			user, _ := ps.db.GetUserByID(ctx, e.Data.ID)
			userEntity := userModelsToEntities([]db.User{user})[0]
			ps.UserPubsub.Publish(userEntity)
		case commands.Operation{Action: commands.CREATE, Entity: commands.TRACK}:
			e, _ := chainutils.FromTxBytes[commands.CreateTrackCommand](tx)
			track, _ := ps.db.GetTrackByID(ctx, e.Data.ID)
			trackEntity := trackModelsToEntities([]db.Track{track})[0]
			ps.TrackPubsub.Publish(trackEntity)
		case commands.Operation{Action: commands.CREATE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.CREATE, Entity: commands.REPOST}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.FOLLOW}:
		case commands.Operation{Action: commands.DELETE, Entity: commands.REPOST}:
		default:
		}
	}
}
