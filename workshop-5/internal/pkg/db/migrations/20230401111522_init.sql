-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                          id BIGSERIAL PRIMARY KEY NOT NULL,
                          name text NOT NULL DEFAULT '',
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
                          updated_at TIMESTAMP WITH TIME ZONE NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
