package graph

import (
	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/graph/model"
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
