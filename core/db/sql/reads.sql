-- name: GetUsers :many
select *
from users
order by created_at;

-- name: GetUserByHandle :one
select *
from users
where handle = $1
limit 1;

-- name: GetUserByID :one
select *
from users
where id = $1
limit 1;

-- name: GetTrackByID :one
select *
from tracks
where id = $1
limit 1;

-- name: GetTracks :many
select *
from tracks
order by created_at;

-- name: GetFollowers :many
select *
from follows
where follower_id = $1
order by created_at;

-- name: GetFollowing :many
select *
from follows
where following_id = $1
order by created_at;

-- name: GetTrackReposts :many
select *
from reposts
where track_id = $1
order by created_at;

-- name: GetUserReposts :many
select *
from reposts
where reposter_id = $1
order by created_at;

-- name: GetTrackByTitle :one
select *
from tracks
where title = $1
limit 1;

-- name: GetFollowByID :one
select *
from follows
where id = $1
limit 1;

-- name: GetRepostByID :one
select *
from reposts
where id = $1
limit 1;