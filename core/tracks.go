package core

import (
	"context"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/core/interface/commands"
	"github.com/alecsavvy/clockwise/core/interface/entities"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
)

func (c *Core) CreateTrack(cmd *commands.CreateTrackCommand) (*entities.TrackEntity, error) {
	ctx := context.Background()
	db := c.db

	// submit command to chain
	_, err := c.Send(cmd)
	if err != nil {
		return nil, err
	}

	track, err := db.GetTrackByTitle(ctx, cmd.Data.Title)
	if err != nil {
		return nil, err
	}

	trackEntity := &entities.TrackEntity{
		ID:          track.ID,
		Title:       track.Title,
		UserID:      track.UserID,
		Description: track.Description,
		Genre:       track.Genre,
		StreamURL:   track.StreamUrl,
	}

	return trackEntity, nil
}

func (c *Core) handleCreateTrack(qtx *db.Queries, createdAt int64, b []byte) (*abcitypes.ExecTxResult, error) {
	ctx := context.Background()

	var cmd commands.CreateTrackCommand
	err := c.fromTxBytes(b, &cmd)
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

func (c *Core) GetTracks() ([]*entities.TrackEntity, error) {
	ctx := context.Background()

	tracks, err := c.db.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	return c.trackModelsToEntities(tracks), nil
}

func (c *Core) GetTrackReposts(trackId string) ([]*entities.RepostEntity, error) {
	ctx := context.Background()
	db := c.db

	reposts, err := db.GetTrackReposts(ctx, trackId)
	if err != nil {
		return nil, utils.AppError("could not get track reposts", err)
	}

	return c.repostModelsToEntities(reposts), nil
}
