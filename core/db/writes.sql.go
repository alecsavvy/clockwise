// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: writes.sql

package db

import (
	"context"
)

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
