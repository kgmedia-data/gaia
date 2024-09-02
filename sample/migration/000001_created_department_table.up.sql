create table if not exists department (
    id bigserial primary key,
    name varchar(255) not null unique,
    is_deleted boolean not null default false,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp
);