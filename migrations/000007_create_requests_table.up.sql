create table requests
(
    id          uuid primary key default gen_random_uuid(),
    description text                          not null,
    geo         geometry(Point, 4326),
    address     text,
    created_at  timestamp                     not null,
    deleted_at  timestamp,
    status      smallint                      not null,
    priority    int,
    theme_id    uuid references themes (id)   not null,
    user_id     uuid references users (id),
    contact_id  uuid references contacts (id) not null
);