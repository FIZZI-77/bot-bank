-- +goose Up
-- +goose StatementBegin
CREATE TYPE currency AS ENUM ('USD', 'EUR', 'RUB');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TYPE currency;
-- +goose StatementEnd
