-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table public.category
(
    id          smallint generated always as identity
        constraint id
            primary key,
    name        varchar(100) not null,
    description varchar(1000),
    "order"     smallint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE public.category;
-- +goose StatementEnd
