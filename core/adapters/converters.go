package adapters

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/utils"
)

func repostModelsToEntities(models []db.Repost) []*entities.RepostEntity {
	return utils.Map(models, func(repost db.Repost) *entities.RepostEntity {
		return &entities.RepostEntity{
			ID:         repost.ID,
			ReposterID: repost.ReposterID,
			TrackID:    repost.TrackID,
		}
	})
}

func trackModelsToEntities(models []db.Track) []*entities.TrackEntity {
	return utils.Map(models, func(track db.Track) *entities.TrackEntity {
		return &entities.TrackEntity{
			ID:          track.ID,
			Title:       track.Title,
			StreamURL:   track.StreamUrl,
			Genre:       track.Genre,
			Description: track.Description,
			UserID:      track.UserID,
		}
	})
}

func userModelsToEntities(models []db.User) []*entities.UserEntity {
	return utils.Map(models, func(user db.User) *entities.UserEntity {
		return &entities.UserEntity{
			ID:      user.ID,
			Handle:  user.Handle,
			Bio:     user.Bio,
			Address: user.Address,
		}
	})
}
