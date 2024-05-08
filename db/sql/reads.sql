-- name: GetUsers :many
select *
from users
order by handle;

-- name: GetTracks :many
select *
from tracks
order by title;

-- name: GetBlocks :many
select *
from blocks
order by blocknumber desc;