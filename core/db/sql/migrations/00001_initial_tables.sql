-- +migrate Up
create table blocks (
    id text primary key,
    blocknumber int not null unique,
    blockhash text not null unique
);

create table users (
    id text primary key,
    handle text not null,
    address text not null,
    bio text not null,
    created_at int not null,
    foreign key (created_at) references blocks (blocknumber) on delete cascade
);

create table tracks (
    id text primary key,
    title text not null,
    stream_url text not null,
    description text not null,
    genre text not null,
    user_id text not null,
    created_at int not null,
    foreign key (user_id) references users (id) on delete cascade,
    foreign key (created_at) references blocks (blocknumber) on delete cascade
);

create table follows (
    id text primary key,
    follower_id text not null,
    following_id text not null,
    created_at int not null,
    unique (follower_id, following_id),
    check (follower_id <> following_id),
    foreign key (follower_id) references users (id) on delete cascade,
    foreign key (following_id) references users (id) on delete cascade,
    foreign key (created_at) references blocks (blocknumber) on delete cascade
);

create table reposts (
    id text primary key,
    reposter_id text not null,
    track_id text not null,
    created_at int not null,
    unique (reposter_id, track_id),
    foreign key (reposter_id) references users(id) on delete cascade,
    foreign key (track_id) references tracks (id) on delete cascade,
    foreign key (created_at) references blocks (blocknumber) on delete cascade
);

-- +migrate Down
drop table if exists reposts;

drop table if exists follows;

drop table if exists tracks;

drop table if exists users;

drop table if exists blocks;