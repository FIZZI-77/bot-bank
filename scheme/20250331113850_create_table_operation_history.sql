-- +goose Up
-- +goose StatementBegin

CREATE TABLE operations(
    id UUID PRIMARY KEY,
    operation_author VARCHAR(255) NOT NULL,
    operation_type VARCHAR(255),
    amount decimal(19,4)  NOT NULL,
    recipient VARCHAR(255)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations CASCADE;
-- +goose StatementEnd