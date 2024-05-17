package handlers

import (
	"context"

	chainutils "github.com/alecsavvy/clockwise/core/chain_utils"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func HandleCreateTrack(qtx *db.Queries, createdAt int32, b []byte) (*abcitypes.ExecTxResult, error) {
	ctx := context.Background()

	cmd, err := chainutils.FromTxBytes[commands.CreateTrackCommand](b)
	if err != nil {
		return nil, utils.AppError("not a create track command in create track handler", err)
	}

	track := cmd.Data

	err = qtx.CreateTrack(ctx, db.CreateTrackParams{
		ID:          track.ID,
		Title:       track.Title,
		Genre:       track.Genre,
		Description: track.Description,
		StreamUrl:   track.StreamURL,
		UserID:      track.UserID,
		CreatedAt:   createdAt,
	})

	if err != nil {
		return nil, utils.AppError("failure to insert track", err)
	}

	return &abcitypes.ExecTxResult{
		Code: 0,
	}, nil
}
