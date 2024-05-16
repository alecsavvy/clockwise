package adapters

import (
	"context"

	chainclient "github.com/alecsavvy/clockwise/core/chain_client"
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/commands"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/cqrs/events"
	"github.com/alecsavvy/clockwise/cqrs/services"
	"github.com/alecsavvy/clockwise/utils"
)

type TrackRepository struct {
	logger *utils.Logger
	cc     *chainclient.ChainClient
	db     *db.Queries
}

// GetTrackReposts implements services.TrackService.
func (t *TrackRepository) GetTrackReposts(trackId string) ([]*entities.RepostEntity, error) {
	ctx := context.Background()
	db := t.db

	reposts, err := db.GetTrackReposts(ctx, trackId)
	if err != nil {
		return nil, utils.AppError("could not get track reposts", err)
	}

	return repostModelsToEntities(reposts), nil
}

// RepostTrack implements services.TrackService.
func (t *TrackRepository) RepostTrack(*commands.Command[commands.CreateRepost]) (*events.RepostCreatedEvent, error) {
	panic("unimplemented")
}

// CreateTrack implements services.TrackService.
func (t *TrackRepository) CreateTrack(*commands.CreateTrackCommand) (*events.TrackCreatedEvent, error) {
	panic("unimplemented")
}

// GetTracks implements services.TrackService.
func (t *TrackRepository) GetTracks() ([]*entities.TrackEntity, error) {
	ctx := context.Background()

	tracks, err := t.db.GetTracks(ctx)
	if err != nil {
		return nil, err
	}

	return trackModelsToEntities(tracks), nil
}

var _ services.TrackService = (*TrackRepository)(nil)
