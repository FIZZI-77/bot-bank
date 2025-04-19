-- +goose Up
-- +goose StatementBegin

CREATE TABLE topup(
    id UUID PRIMARY KEY,
    user_telegram_id BIGINT NOT NULL,
    amount decimal(18,2)  NOT NULL,
    currency currency NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE topup CASCADE;
-- +goose StatementEnd