-- +goose Up
-- +goose StatementBegin

CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    sender_uuid UUID NOT NULL,
    amount decimal(18,2)  NOT NULL,
    currency currency NOT NULL ,
    recipient_tg_id BIGINT

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions CASCADE;
-- +goose StatementEnd