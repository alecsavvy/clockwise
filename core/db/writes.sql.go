// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: writes.sql

package db

import (
	"context"
)

const createFollow = `-- name: CreateFollow :exec
insert into follows (follower_id, following_id)
values ($1, $2)
`

type CreateFollowParams struct {
	FollowerID  string
	FollowingID string
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.db.Exec(ctx, createFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const createRepost = `-- name: CreateRepost :exec
insert into reposts (track_id, reposter_id)
values ($1, $2)
`

type CreateRepostParams struct {
	TrackID    string
	ReposterID string
}

func (q *Queries) CreateRepost(ctx context.Context, arg CreateRepostParams) error {
	_, err := q.db.Exec(ctx, createRepost, arg.TrackID, arg.ReposterID)
	return err
}

const createTrack = `-- name: CreateTrack :exec
insert into tracks (id, title, stream_url, description, user_id)
values ($1, $2, $3, $4, $5)
`

type CreateTrackParams struct {
	ID          string
	Title       string
	StreamUrl   string
	Description string
	UserID      string
}

func (q *Queries) CreateTrack(ctx context.Context, arg CreateTrackParams) error {
	_, err := q.db.Exec(ctx, createTrack,
		arg.ID,
		arg.Title,
		arg.StreamUrl,
		arg.Description,
		arg.UserID,
	)
	return err
}

const createUser = `-- name: CreateUser :exec
insert into users (id, handle, address, bio)
values ($1, $2, $3, $4)
`

type CreateUserParams struct {
	ID      string
	Handle  string
	Address string
	Bio     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.ID,
		arg.Handle,
		arg.Address,
		arg.Bio,
	)
	return err
}

const removeFollow = `-- name: RemoveFollow :exec
delete from follows
where follower_id = $1
  and following_id = $2
`

type RemoveFollowParams struct {
	FollowerID  string
	FollowingID string
}

func (q *Queries) RemoveFollow(ctx context.Context, arg RemoveFollowParams) error {
	_, err := q.db.Exec(ctx, removeFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const removeRepost = `-- name: RemoveRepost :exec
delete from reposts
where reposter_id = $1
  and track_id = $2
`

type RemoveRepostParams struct {
	ReposterID string
	TrackID    string
}

func (q *Queries) RemoveRepost(ctx context.Context, arg RemoveRepostParams) error {
	_, err := q.db.Exec(ctx, removeRepost, arg.ReposterID, arg.TrackID)
	return err
}
