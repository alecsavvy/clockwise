// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: reads.sql

package db

import (
	"context"
)

const getTrack = `-- name: GetTrack :one
select id, title, stream_url, description, user_id, tx_hash
from tracks
where id = $1
`

func (q *Queries) GetTrack(ctx context.Context, id string) (Track, error) {
	row := q.db.QueryRow(ctx, getTrack, id)
	var i Track
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.StreamUrl,
		&i.Description,
		&i.UserID,
		&i.TxHash,
	)
	return i, err
}

const getTrackReposts = `-- name: GetTrackReposts :many
select reposter_id, track_id, tx_hash
from reposts
where track_id = $1
`

func (q *Queries) GetTrackReposts(ctx context.Context, trackID string) ([]Repost, error) {
	rows, err := q.db.Query(ctx, getTrackReposts, trackID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repost
	for rows.Next() {
		var i Repost
		if err := rows.Scan(&i.ReposterID, &i.TrackID, &i.TxHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTracks = `-- name: GetTracks :many
select id, title, stream_url, description, user_id, tx_hash
from tracks
order by title
`

func (q *Queries) GetTracks(ctx context.Context) ([]Track, error) {
	rows, err := q.db.Query(ctx, getTracks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.StreamUrl,
			&i.Description,
			&i.UserID,
			&i.TxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTxResult = `-- name: GetTxResult :one
select rowid, block_id, index, created_at, tx_hash, tx_result from tx_results where tx_hash = $1
`

func (q *Queries) GetTxResult(ctx context.Context, txHash string) (TxResult, error) {
	row := q.db.QueryRow(ctx, getTxResult, txHash)
	var i TxResult
	err := row.Scan(
		&i.Rowid,
		&i.BlockID,
		&i.Index,
		&i.CreatedAt,
		&i.TxHash,
		&i.TxResult,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
select id, handle, bio, tx_hash
from users
where id = $1
limit 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Handle,
		&i.Bio,
		&i.TxHash,
	)
	return i, err
}

const getUserByHandle = `-- name: GetUserByHandle :one
select id, handle, bio, tx_hash
from users
where handle = $1
limit 1
`

func (q *Queries) GetUserByHandle(ctx context.Context, handle string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByHandle, handle)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Handle,
		&i.Bio,
		&i.TxHash,
	)
	return i, err
}

const getUserFollowers = `-- name: GetUserFollowers :many
select follower_id, following_id, tx_hash
from follows
where following_id = $1
`

func (q *Queries) GetUserFollowers(ctx context.Context, followingID string) ([]Follow, error) {
	rows, err := q.db.Query(ctx, getUserFollowers, followingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowerID, &i.FollowingID, &i.TxHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserFollowing = `-- name: GetUserFollowing :many
select follower_id, following_id, tx_hash
from follows
where follower_id = $1
`

func (q *Queries) GetUserFollowing(ctx context.Context, followerID string) ([]Follow, error) {
	rows, err := q.db.Query(ctx, getUserFollowing, followerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(&i.FollowerID, &i.FollowingID, &i.TxHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserReposts = `-- name: GetUserReposts :many
select reposter_id, track_id, tx_hash
from reposts
where reposter_id = $1
`

func (q *Queries) GetUserReposts(ctx context.Context, reposterID string) ([]Repost, error) {
	rows, err := q.db.Query(ctx, getUserReposts, reposterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Repost
	for rows.Next() {
		var i Repost
		if err := rows.Scan(&i.ReposterID, &i.TrackID, &i.TxHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTracks = `-- name: GetUserTracks :many
select id, title, stream_url, description, user_id, tx_hash
from tracks
where user_id = $1
`

func (q *Queries) GetUserTracks(ctx context.Context, userID string) ([]Track, error) {
	rows, err := q.db.Query(ctx, getUserTracks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.StreamUrl,
			&i.Description,
			&i.UserID,
			&i.TxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsers = `-- name: GetUsers :many
select id, handle, bio, tx_hash
from users
order by handle
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Handle,
			&i.Bio,
			&i.TxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
