create table contacts
(
    id    uuid primary key default gen_random_uuid(),
    phone text,
    email text,
    note  text,
    name  text,
    deleted_at  timestamp
);