package handlers

import (
	"context"
	"fmt"

	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func HandleCreateUser(qtx *db.Queries, createdAt int32, b []byte) (*abcitypes.ExecTxResult, error) {
	ctx := context.Background()

	cmd, err := chainutils.FromTxBytes[commands.CreateUserCommand](b)
	if err != nil {
		return nil, utils.AppError("not a create user command in create user handler", err)
	}

	user := cmd.Data

	err = qtx.CreateUser(ctx, db.CreateUserParams{
		ID:        user.ID,
		Handle:    user.Handle,
		Bio:       user.Bio,
		Address:   user.Address,
		CreatedAt: createdAt,
	})

	if err != nil {
		return nil, utils.AppError("failure to insert user", err)
	}

	eventAttributes := make([]abcitypes.EventAttribute, 1)
	eventAttributes = append(eventAttributes, abcitypes.EventAttribute{
		Key:   user.ID,
		Value: user.Handle,
	})

	createUserEvent := abcitypes.Event{
		Type:       fmt.Sprintf("%s%s", cmd.Action, cmd.Entity),
		Attributes: eventAttributes,
	}
	events := make([]abcitypes.Event, 1)
	events = append(events, createUserEvent)

	return &abcitypes.ExecTxResult{
		Code:   0,
		Events: events,
	}, nil
}
