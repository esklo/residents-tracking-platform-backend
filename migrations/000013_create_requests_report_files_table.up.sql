create table requests_report_files
(
    request_id uuid references requests (id) not null,
    file_id    uuid references files (id)    not null
);