package core

import (
	"context"

	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/utils"
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
