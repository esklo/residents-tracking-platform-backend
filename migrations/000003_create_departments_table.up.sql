create table departments
(
    id          uuid primary key default gen_random_uuid(),
    title       text                           not null,
    district_id uuid references districts (id) not null,
    deleted_at  timestamp
);