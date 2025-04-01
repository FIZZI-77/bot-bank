-- +goose Up
-- +goose StatementBegin

CREATE TABLE operations(
    id serial PRIMARY KEY,
    operation_author VARCHAR(255) REFERENCES users(username) ON DELETE CASCADE NOT NULL,
    operation VARCHAR(255),
    amount int NOT NULL,
    recipient VARCHAR(255)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations CASCADE;
-- +goose StatementEnd