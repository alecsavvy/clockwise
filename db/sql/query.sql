-- name: ListUsers :many
select *
from users
order by handle;

-- name: ListTracks :many
select *
from tracks
order by title;