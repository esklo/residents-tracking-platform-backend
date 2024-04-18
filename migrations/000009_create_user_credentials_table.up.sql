create table user_credentials
(
    user_id        uuid references users (id) not null,
    credential_id  bytea                      not null unique,
    credential     bytea                      not null,
    created_at     timestamp,
    last_used_at   timestamp,
    requested_from text
);