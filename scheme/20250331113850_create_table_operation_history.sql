-- +goose Up
-- +goose StatementBegin

CREATE TABLE operations_history(
    id UUID PRIMARY KEY,
    sender UUID NOT NULL,
    operation_type VARCHAR(255),
    amount decimal(19,4)  NOT NULL,
    recipient UUID

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations CASCADE;
-- +goose StatementEnd