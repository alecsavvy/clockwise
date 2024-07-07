-- name: CreateUser :exec
insert into users (id, handle, bio)
values ($1, $2, $3);

-- name: CreateTrack :exec
insert into tracks (id, title, stream_url, description, user_id)
values ($1, $2, $3, $4, $5);

-- name: CreateFollow :exec
insert into follows (follower_id, following_id)
values ($1, $2);

-- name: CreateRepost :exec
insert into reposts (track_id, reposter_id)
values ($1, $2);

-- name: RemoveFollow :exec
delete from follows
where follower_id = $1
  and following_id = $2;

-- name: RemoveRepost :exec
delete from reposts
where reposter_id = $1
  and track_id = $2;