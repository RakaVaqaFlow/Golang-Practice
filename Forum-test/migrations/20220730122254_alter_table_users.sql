-- +goose Up
-- +goose StatementBegin
AFTER TABLE users
    ADD COLUMN IF NOT EXISTS avatar text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
AFTER TABLE users
    DROP COLUMN IF EXISTS avatar text;
-- +goose StatementEnd
