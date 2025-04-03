-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id UUID PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users CASCADE;
-- +goose StatementEnd