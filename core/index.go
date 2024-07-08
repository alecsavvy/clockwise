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
		txStatusCode := abcitypes.CodeTypeOK
		if err := protocol.MessageRouter(ctx, c.indexingRoutes, tx); err != nil {
			// if certain errors not ok code, in others block consensus (halt)
			txStatusCode = CodeTypeNotOK
		}
		txResults[i] = &abcitypes.ExecTxResult{
			Code: txStatusCode,
		}
	}

	return txResults, nil
}

func (c *Core) indexCreateUser(ctx context.Context, msg proto.Message) error {
	logger := c.logger
	qtx := c.getDb()

	tx, ok := msg.(*gen.CreateUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
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
		return errors.New("invalid msg passed to validator")
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
	return errors.New("unimplemented")
}

func (c *Core) indexUnfollowUser(ctx context.Context, msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexRepostTrack(ctx context.Context, msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexUnrepostTrack(ctx context.Context, msg proto.Message) error {
	return errors.New("unimplemented")
}
