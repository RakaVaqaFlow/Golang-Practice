-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users (
    nickname  varchar(255) NOT NULL CONSTRAINT nick_right CHECK(nickname ~ '^[A-Za-z0-9_\.]*$') PRIMARY KEY,
    email     varchar(255) NOT NULL UNIQUE CONSTRAINT email_right CHECK(email ~ '^.*@[A-Za-z0-9\-_\.]*$'),
    full_name varchar(255) NOT NULL,
    about     text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
