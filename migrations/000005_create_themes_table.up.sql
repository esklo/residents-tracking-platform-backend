create table themes
(
    id            uuid primary key default gen_random_uuid(),
    title         text                             not null,
    priority      int                              not null,
    department_id uuid references departments (id) not null,
    deleted_at  timestamp
);