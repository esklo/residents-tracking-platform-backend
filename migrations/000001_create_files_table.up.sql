create table files
(
    id   uuid primary key default gen_random_uuid(),
    filename text not null,
    mimetype text not null,
    extension text not null,
    path text not null,
    deleted_at  timestamp
);