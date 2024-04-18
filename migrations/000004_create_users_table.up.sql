create table users
(
    id            uuid primary key default gen_random_uuid(),
    email         text not null,
    role          text not null,
    first_name    text not null,
    last_name     text,
    father_name   text,
    password      text not null,
    salt          text not null,
    department_id uuid references departments (id),
    deleted_at  timestamp
);