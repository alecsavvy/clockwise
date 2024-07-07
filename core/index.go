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
		if err := protocol.MessageRouter(c.indexingRoutes, tx); err != nil {
			// if certain errors not ok code, in others block consensus (halt)
			txStatusCode = CodeTypeNotOK
		}
		txResults[i] = &abcitypes.ExecTxResult{
			Code: txStatusCode,
		}
	}

	return txResults, nil
}

func (c *Core) indexCreateUser(msg proto.Message) error {
	ctx := context.Background()
	logger := c.logger
	qtx := c.getDb()

	tx, ok := msg.(*gen.CreateUser)
	if !ok {
		return errors.New("invalid msg passed to validator")
	}

	data := tx.GetData()

	args := db.CreateUserParams{
		ID:      data.Id,
		Handle:  data.Handle,
		Address: data.Address,
		Bio:     data.Bio,
	}

	if err := qtx.CreateUser(ctx, args); err != nil {
		logger.Error("error persisting new user", "user", args, "error", err)
		return err
	}

	return nil
}

func (c *Core) indexFollowUser(msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexUnfollowUser(msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexCreateTrack(msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexRepostTrack(msg proto.Message) error {
	return errors.New("unimplemented")
}

func (c *Core) indexUnrepostTrack(msg proto.Message) error {
	return errors.New("unimplemented")
}
