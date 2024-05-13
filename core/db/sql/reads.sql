-- name: GetUsers :many
select *
from users
order by handle;

-- name: GetUserByHandle :one
select *
from users
where handle = $1
limit 1;

-- name: GetTracks :many
select *
from tracks
order by title;

-- name: GetFollowersByHandle :many
select u2.*
from users u1
    join follows f on u1.id = f.following_id
    join users u2 on f.follower_id = u2.id
where u1.handle = $1;