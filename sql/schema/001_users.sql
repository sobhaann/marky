-- +goose up
create table users(
    id bigint primary key generated always as identity,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp default null,
    r_id uuid not null unique default gen_random_uuid(),
    email text unique not null
);

-- +goose down
DROP TABLE users;
