-- +migrate Up
create table users (
    id text primary key,
    handle text not null,
    address text not null,
    bio text not null
);

create table tracks (
    id text primary key,
    title text not null,
    stream_url text not null,
    description text not null,
    user_id text not null,
    foreign key (user_id) references users (id) on delete cascade
);

create table follows (
    follower_id text references users(id),
    following_id text references users(id),
    primary key (follower_id, following_id),
    check (follower_id <> following_id)
);

create table reposts (
    reposter_id text references users(id),
    track_id text references tracks(id),
    primary key (reposter_id, track_id)
);

-- +migrate Down
drop table if exists users;

drop table if exists tracks;

drop table if exists follows;

drop table if exists reposts;