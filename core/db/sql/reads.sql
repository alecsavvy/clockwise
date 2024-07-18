-- name: GetUsers :many
select *
from users
order by handle;

-- name: GetTracks :many
select *
from tracks
order by title;

-- name: GetUserByHandle :one
select *
from users
where handle = $1
limit 1;

-- name: GetUser :one
select *
from users
where id = $1
limit 1;

-- name: GetUserTracks :many
select *
from tracks
where user_id = $1;

-- name: GetUserFollowers :many
select *
from follows
where following_id = $1;

-- name: GetUserFollowing :many
select *
from follows
where follower_id = $1;

-- name: GetUserReposts :many
select *
from reposts
where reposter_id = $1;

-- name: GetTrackReposts :many
select *
from reposts
where track_id = $1;

-- name: GetTrack :one
select *
from tracks
where id = $1;

-- name: GetTxResult :one
select * from tx_results where tx_hash = $1;
