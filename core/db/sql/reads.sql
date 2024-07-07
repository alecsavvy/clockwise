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

-- name: GetFollowersByHandle :many
select u2.*
from users u1
    join follows f on u1.id = f.following_id
    join users u2 on f.follower_id = u2.id
where u1.handle = $1;