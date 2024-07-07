package db

import "context"

type UserData struct {
	User
	Tracks    []Track
	Following []Follow
	Followers []Follow
	Reposts   []Repost
}

func (q *Queries) GetUserData(ctx context.Context, address string) (*UserData, error) {
	user, err := q.GetUser(ctx, address)
	if err != nil {
		return nil, err
	}

	tracks, err := q.GetUserTracks(ctx, address)
	if err != nil {
		return nil, err
	}

	following, err := q.GetUserFollowing(ctx, address)
	if err != nil {
		return nil, err
	}

	followers, err := q.GetUserFollowers(ctx, address)
	if err != nil {
		return nil, err
	}

	reposts, err := q.GetUserReposts(ctx, address)
	if err != nil {
		return nil, err
	}

	return &UserData{
		User:      user,
		Tracks:    tracks,
		Following: following,
		Followers: followers,
		Reposts:   reposts,
	}, nil
}
