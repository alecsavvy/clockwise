package graph

import (
	"github.com/alecsavvy/clockwise/core/interface/entities"
	"github.com/alecsavvy/clockwise/ports/graph/model"
	"github.com/alecsavvy/clockwise/utils"
)

func (r *Resolver) followEntitiesToModels(models []*entities.FollowEntity) []*model.Follow {
	return utils.Map(models, func(follow *entities.FollowEntity) *model.Follow {
		return &model.Follow{
			ID:          follow.ID,
			FollowerID:  follow.FollowerID,
			FollowingID: follow.FollowingID,
		}
	})
}

func (r *Resolver) userEntitiesToModels(models []*entities.UserEntity) []*model.User {
	return utils.Map(models, func(entity *entities.UserEntity) *model.User {
		return &model.User{
			ID:      entity.ID,
			Handle:  entity.Handle,
			Bio:     entity.Bio,
			Address: entity.Address,
		}
	})
}
