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

	trackEntities := utils.Map(tracks, func(track db.Track) *entities.TrackEntity {
		return &entities.TrackEntity{
			ID:          track.ID,
			Title:       track.Title,
			StreamURL:   track.StreamUrl,
			Description: track.Description,
			UserID:      track.UserID,
		}
	})

	return trackEntities, nil
}

var _ services.TrackService = (*TrackRepository)(nil)
