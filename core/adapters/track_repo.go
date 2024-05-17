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

// CreateTrackEvents implements services.TrackService.
func (t *TrackRepository) CreateTrackEvents() (<-chan *entities.TrackEntity, error) {
	panic("unimplemented")
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
func (t *TrackRepository) CreateTrack(cmd *commands.CreateTrackCommand) (*events.TrackCreatedEvent, error) {
	ctx := context.Background()
	cc := t.cc
	db := t.db

	// submit command to chain
	res, err := cc.Send(cmd)
	if err != nil {
		return nil, err
	}

	// construct event
	var event events.TrackCreatedEvent
	event.BlockHeight = uint64(res.Height)
	event.TransactionHash = string(res.Hash)

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

	event.Track = *trackEntity

	return &event, nil
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

func NewTrackRepo(logger *utils.Logger, cc *chainclient.ChainClient, db *db.Queries) *TrackRepository {
	return &TrackRepository{
		logger: logger,
		cc:     cc,
		db:     db,
	}
}

var _ services.TrackService = (*TrackRepository)(nil)
