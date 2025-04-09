-- +goose Up
-- +goose StatementBegin

CREATE TABLE transactions_history(
    id SERIAL PRIMARY KEY,
    sender BIGINT NOT NULL,
    amount decimal(18,2)  NOT NULL,
    recipient BIGINT

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions_history CASCADE;
-- +goose StatementEnd