-- +goose Up
-- +goose StatementBegin

CREATE TABLE topup_history(
    id SERIAL PRIMARY KEY,
    user_telegram_id BIGINT NOT NULL,
    amount decimal(18,2)  NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE topup_history CASCADE;
-- +goose StatementEnd