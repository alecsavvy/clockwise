-- +migrate Up
create table users (
    id uuid primary key,
    handle text not null,
    address text not null,
    bio text
);

create table tracks (
    id uuid primary key,
    title text not null,
    stream_url text not null,
    description text,
    user_id uuid not null,
    foreign key (user_id) references users (id) on delete cascade
);

create table follows (
    follower_id uuid references users(id),
    following_id uuid references users(id),
    primary key (follower_id, following_id),
    check (follower_id <> following_id)
);

create table reposts (
    reposter_id uuid references users(id),
    track_id uuid references tracks(id),
    primary key (reposter_id, track_id)
);

-- +migrate Down
drop table if exists users;

drop table if exists tracks;

drop table if exists follows;

drop table if exists reposts;