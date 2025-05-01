package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"tgtransaction/src/core/models"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (c *TransactionPostgres) PersistTransaction(ctx context.Context, model models.TransactionModel) error {
	var currency string
	currency = FromCurrencyEnum(model.Cur)

	const sendQuery = "INSERT INTO transactions (id, sender_id, amount, currency, recipient_tg_id) VALUES ($1, $2, $3, $4, $5)"

	_, err := c.db.ExecContext(ctx, sendQuery, model.TransactionId, model.UserId, model.Amount, currency, model.Recipient)
	if err != nil {

		return fmt.Errorf("run sql sendQuery: %w", err)
	}
	return nil
}

func (c *TransactionPostgres) GetTotalSentAmount(ctx context.Context, userID uuid.UUID) (_ models.Balance, err error) {

	const getTotalSendQuery = `SELECT
    	currency,
		COALESCE(SUM(CASE WHEN sender_id = $1 THEN amount ELSE 0 END ),0)
	FROM transactions
	GROUP BY currency
`
	rows, err := c.db.QueryContext(ctx, getTotalSendQuery, userID)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return models.Balance{}, fmt.Errorf("getTotalSendQuery timed out: %w", err)
		}
		return models.Balance{}, fmt.Errorf("run sql getTotalSendQuery: %w", err)
	}

	defer func() {
		if closeErr := rows.Close(); err != nil {
			err = errors.Join(err, fmt.Errorf("close rows %w", closeErr))
		}
	}()

	balance := models.Balance{}
	for rows.Next() {
		var currency string
		var amount decimal.Decimal
		if err := rows.Scan(&currency, &amount); err != nil {
			return models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		balance[curEnum] = amount
	}
	return balance, nil
}

func (c *TransactionPostgres) GetTotalReceivedAmount(ctx context.Context, userTgID int64) (_ models.Balance, err error) {

	const getTotalReceivedQuery = `SELECT
    currency,
		COALESCE(SUM(CASE WHEN recipient_tg_id = $1 THEN amount ELSE 0 END ),0)
	FROM transactions
	GROUP BY currency
`
	rows, err := c.db.QueryContext(ctx, getTotalReceivedQuery, userTgID)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return models.Balance{}, fmt.Errorf("getTotalReceivedQuery timed out: %w", err)
		}
		return models.Balance{}, fmt.Errorf("run sql getTotalSendQuery: %w", err)
	}

	defer func() {
		if closeErr := rows.Close(); err != nil {
			err = errors.Join(err, fmt.Errorf("close rows %w", closeErr))
		}
	}()

	balance := models.Balance{}
	for rows.Next() {
		var currency string
		var amount decimal.Decimal
		if err := rows.Scan(&currency, &amount); err != nil {
			return models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		balance[curEnum] = amount
	}
	return balance, nil
}
