package graph

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/graph/model"
	"github.com/alecsavvy/clockwise/protocol/gen"
	"github.com/alecsavvy/clockwise/utils"
)

func dbUserDataToUserModel(user *db.UserData) *model.User {
	return &model.User{
		Address:   user.ID,
		Handle:    user.Handle,
		Bio:       user.Bio,
		Tracks:    dbTrackToTrackModel(user.Tracks),
		Followers: dbFollowToFollowModel(user.Followers),
		Following: dbFollowToFollowModel(user.Following),
		Reposts:   dbRepostToRepostModel(user.Reposts),
	}
}

func dbUserToUserModel(users []db.User) []*model.User {
	return utils.Map(users, func(user db.User) *model.User {
		return &model.User{
			Address: user.ID,
			Handle:  user.Handle,
			Bio:     user.Bio,
		}
	})
}

func dbTrackToTrackModel(tracks []db.Track) []*model.Track {
	return utils.Map(tracks, func(t db.Track) *model.Track {
		return &model.Track{
			ID:          t.ID,
			Title:       t.Title,
			StreamURL:   t.StreamUrl,
			Description: t.Description,
			UserID:      t.UserID,
		}
	})
}

func dbFollowToFollowModel(follows []db.Follow) []*model.Follow {
	return utils.Map(follows, func(f db.Follow) *model.Follow {
		return &model.Follow{
			FollowerID: f.FollowerID,
			FolloweeID: f.FollowingID,
		}
	})
}

func dbRepostToRepostModel(reposts []db.Repost) []*model.Repost {
	return utils.Map(reposts, func(r db.Repost) *model.Repost {
		return &model.Repost{
			ReposterID: r.ReposterID,
			TrackID:    r.TrackID,
		}
	})
}

func protoToTrackModel(tracks []*gen.CreateTrack) []*model.Track {
	return utils.Map(tracks, func(t *gen.CreateTrack) *model.Track {
		return &model.Track{
			ID:          t.Data.Id,
			Title:       t.Data.Title,
			StreamURL:   t.Data.StreamUrl,
			Description: t.Data.Description,
			UserID:      t.Data.UserId,
		}
	})
}

func protoToUserModel(users []*gen.CreateUser) []*model.User {
	return utils.Map(users, func(user *gen.CreateUser) *model.User {
		return &model.User{
			Address: user.Data.Address,
			Handle:  user.Data.Handle,
			Bio:     user.Data.Bio,
		}
	})
}

func protoToFollowModel(follows []*gen.FollowUser) []*model.Follow {
	return utils.Map(follows, func(f *gen.FollowUser) *model.Follow {
		return &model.Follow{
			FollowerID: f.Data.FollowerId,
			FolloweeID: f.Data.FolloweeId,
		}
	})
}

func protoToRepostModel(reposts []*gen.RepostTrack) []*model.Repost {
	return utils.Map(reposts, func(r *gen.RepostTrack) *model.Repost {
		return &model.Repost{
			ReposterID: r.Data.ReposterId,
			TrackID:    r.Data.TrackId,
		}
	})
}
