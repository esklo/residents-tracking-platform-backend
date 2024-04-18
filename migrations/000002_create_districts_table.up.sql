create table districts
(
    id                   uuid primary key default gen_random_uuid(),
    title                text not null,
    geojson              bytea,
    coat_of_arms_file_id uuid references files (id),
    deleted_at  timestamp
);