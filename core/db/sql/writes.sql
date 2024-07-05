-- name: CreateUser :exec
insert into users (id, handle, address, bio)
values ($1, $2, $3, $4);

-- name: CreateTrack :exec
insert into tracks (id, title, stream_url, description, user_id)
values ($1, $2, $3, $4, $5);