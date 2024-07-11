package core

import (
	"context"
	"errors"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/protocol"
	"github.com/alecsavvy/clockwise/protocol/gen"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"google.golang.org/protobuf/proto"
)

func (c *Core) indexTxs(ctx context.Context, rfb *abcitypes.RequestFinalizeBlock) ([]*abcitypes.ExecTxResult, error) {
	c.startInProgressTx(ctx)

	var txResults = make([]*abcitypes.ExecTxResult, len(rfb.Txs))
	for i, tx := range rfb.Txs {
		// check validation again before indexing
		if err := protocol.MessageRouter(ctx, c.validationRoutes, tx); err != nil {
			txResults[i] = &abcitypes.ExecTxResult{
				Code: CodeTypeNotOK,
				Log: err.Error(),
			}
			continue
		}

		// index transaction after secondary validation
		if err := protocol.MessageRouter(ctx, c.indexingRoutes, tx); err != nil {
			// if certain errors not ok code, in others block consensus (halt)
			txResults[i] = &abcitypes.ExecTxResult{
				Code: CodeTypeNotOK,
				Log: err.Error(),
			}
			continue
		}
		txResults[i] = &abcitypes.ExecTxResult{
			Code: abcitypes.CodeTypeOK,
		}
	}

	return txResults, nil
}

func (c *Core) indexCreateUser(ctx context.Context, msg proto.Message) error {
	logger := c.logger
	qtx := c.getDb()

	tx, ok := msg.(*gen.CreateUser)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	data := tx.GetData()

	args := db.CreateUserParams{
		ID:     data.Address,
		Handle: data.Handle,
		Bio:    data.Bio,
	}

	if err := qtx.CreateUser(ctx, args); err != nil {
		logger.Error("error persisting new user", "user", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexCreateTrack(ctx context.Context, msg proto.Message) error {
	logger := c.logger
	qtx := c.getDb()

	tx, ok := msg.(*gen.CreateTrack)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	data := tx.GetData()

	args := db.CreateTrackParams{
		ID:          data.Id,
		Title:       data.Title,
		StreamUrl:   data.StreamUrl,
		Description: data.Description,
		UserID:      data.UserId,
	}

	if err := qtx.CreateTrack(ctx, args); err != nil {
		logger.Error("error persisting new track", "track", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexFollowUser(ctx context.Context, msg proto.Message) error {
	tx, ok := msg.(*gen.FollowUser)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	qtx := c.getDb()

	args := db.CreateFollowParams{
		FollowerID:  tx.Data.FollowerId,
		FollowingID: tx.Data.FolloweeId,
	}

	if err := qtx.CreateFollow(ctx, args); err != nil {
		c.logger.Error("error persisting new follow", "follow", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexUnfollowUser(ctx context.Context, msg proto.Message) error {
	tx, ok := msg.(*gen.UnfollowUser)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	qtx := c.getDb()

	args := db.RemoveFollowParams{
		FollowerID:  tx.Data.FollowerId,
		FollowingID: tx.Data.FolloweeId,
	}

	if err := qtx.RemoveFollow(ctx, args); err != nil {
		c.logger.Error("error removing follow", "follow", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexRepostTrack(ctx context.Context, msg proto.Message) error {
	tx, ok := msg.(*gen.RepostTrack)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	qtx := c.getDb()

	args := db.CreateRepostParams{
		ReposterID: tx.Data.ReposterId,
		TrackID:    tx.Data.TrackId,
	}

	if err := qtx.CreateRepost(ctx, args); err != nil {
		c.logger.Error("error persisting new repost", "repost", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexUnrepostTrack(ctx context.Context, msg proto.Message) error {
	tx, ok := msg.(*gen.UnrepostTrack)
	if !ok {
		return errors.New("invalid msg passed to indexer")
	}

	qtx := c.getDb()

	args := db.RemoveRepostParams{
		ReposterID: tx.Data.ReposterId,
		TrackID:    tx.Data.TrackId,
	}

	if err := qtx.RemoveRepost(ctx, args); err != nil {
		c.logger.Error("error removing repost", "repost", args, "error", err)
		return err
	}

	return nil
}
