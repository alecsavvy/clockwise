package client

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/cqrs/entities"
	"github.com/alecsavvy/clockwise/utils"
)

func (c *Core) repostModelsToEntities(models []db.Repost) []*entities.RepostEntity {
	return utils.Map(models, func(repost db.Repost) *entities.RepostEntity {
		return &entities.RepostEntity{
			ID:         repost.ID,
			ReposterID: repost.ReposterID,
			TrackID:    repost.TrackID,
		}
	})
}

func (c *Core) trackModelsToEntities(models []db.Track) []*entities.TrackEntity {
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

func (c *Core) userModelsToEntities(models []db.User) []*entities.UserEntity {
	return utils.Map(models, func(user db.User) *entities.UserEntity {
		return &entities.UserEntity{
			ID:      user.ID,
			Handle:  user.Handle,
			Bio:     user.Bio,
			Address: user.Address,
		}
	})
}

func (c *Core) followModelsToEntities(models []db.Follow) []*entities.FollowEntity {
	return utils.Map(models, func(follow db.Follow) *entities.FollowEntity {
		return &entities.FollowEntity{
			ID:          follow.ID,
			FollowerID:  follow.FollowerID,
			FollowingID: follow.FollowingID,
		}
	})
}
