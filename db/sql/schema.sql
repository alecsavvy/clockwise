create table users (
    id uuid primary key,
    handle text not null,
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

create table blocks (
    blocknumber bigint primary key,
    blocktime date not null
);