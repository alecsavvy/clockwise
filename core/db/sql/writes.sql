-- name: CreateUser :exec
insert into users (id, handle, address, bio)
values ($1, $2, $3, $4);

-- name: CreateTrack :exec
insert into tracks (
        id,
        title,
        genre,
        stream_url,
        description,
        user_id
    )
values ($1, $2, $3, $4, $5, $6);

-- name: CreateFollow :exec
insert into follows (id, follower_id, following_id)
values ($1, $2, $3);

-- name: CreateRepost :exec
insert into reposts (id, reposter_id, track_id)
values ($1, $2, $3);

-- name: DeleteFollow :exec
delete from follows
where id = $1;

-- name: DeleteRepost :exec
delete from reposts
where id = $1;

-- name: CreateBlock :exec
insert into blocks (id, blocknumber, blockhash)
values ($1, $2, $3);