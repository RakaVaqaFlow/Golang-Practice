-- +goose Up
-- +goose StatementBegin
create table if not exists todo (
    id bigserial,
    user_id text not null ,
    text text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
