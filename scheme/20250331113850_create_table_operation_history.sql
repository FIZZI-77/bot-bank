-- +goose Up
-- +goose StatementBegin

CREATE TABLE operations_history(
    id SERIAL PRIMARY KEY,
    sender BIGINT NOT NULL,
    operation_type VARCHAR(255),
    amount decimal(18,2)  NOT NULL,
    recipient BIGINT

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations_history CASCADE;
-- +goose StatementEnd