-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id serial PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users CASCADE;
-- +goose StatementEnd