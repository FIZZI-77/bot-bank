-- +goose Up
-- +goose StatementBegin

CREATE TABLE top_up(
    id UUID PRIMARY KEY,
    user_telegram_id BIGINT NOT NULL,
    amount decimal(18,2)  NOT NULL,
    currency currency NOT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE top_up CASCADE;
-- +goose StatementEnd