// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: writes.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTrack = `-- name: CreateTrack :exec
insert into tracks (id, title, stream_url, description, user_id)
values ($1, $2, $3, $4, $5)
`

type CreateTrackParams struct {
	ID          pgtype.UUID
	Title       string
	StreamUrl   string
	Description pgtype.Text
	UserID      pgtype.UUID
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
insert into users (id, handle, bio)
values ($1, $2, $3)
`

type CreateUserParams struct {
	ID     pgtype.UUID
	Handle string
	Bio    pgtype.Text
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.ID, arg.Handle, arg.Bio)
	return err
}

const insertBlock = `-- name: InsertBlock :exec
insert into blocks (blocknumber, blocktime)
values ($1, $2)
`

type InsertBlockParams struct {
	Blocknumber int64
	Blocktime   pgtype.Date
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) error {
	_, err := q.db.Exec(ctx, insertBlock, arg.Blocknumber, arg.Blocktime)
	return err
}