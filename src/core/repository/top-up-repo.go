package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"tgtransaction/src/core/models"
)

type TopUpPostgres struct {
	db *sql.DB
}

func NewTopUpPostgres(db *sql.DB) *TopUpPostgres {
	return &TopUpPostgres{db: db}
}

func (c *TopUpPostgres) PersistTopUp(model *models.TopUpModel) error {
	var currency string
	currency = FromCurrencyEnum(model.Cur)

	const topUpQuery = "INSERT INTO top_up (id, user_telegram_id, amount, currency) VALUES ($1, $2, $3, $4)"
	_, err := c.db.Exec(topUpQuery, model.TopUpId, model.UserId, model.Amount, currency)
	if err != nil {
		return fmt.Errorf("run sql topUpQuery: %w", err)
	}
	return nil
}

func (c *TopUpPostgres) GetTotalTopupAmount(ctx context.Context, userTgID int64) (_ *models.Balance, err error) {
	totalTopupAmount := &models.Balance{}
	if totalTopupAmount == nil {
		return &models.Balance{}, fmt.Errorf("totalTopupAmount nil ptr")
	}

	const totalTopUpAmountQuery = `SELECT
    currency,
	COALESCE(SUM(CASE WHEN user_telegram_id = $1 THEN amount ELSE 0 END),0)
	FROM top_up
	GROUP BY currency`

	rows, err := c.db.QueryContext(ctx, totalTopUpAmountQuery, userTgID)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return &models.Balance{}, fmt.Errorf("totalTopUpAmountQuery timed out: %w", err)
		}
		return &models.Balance{}, fmt.Errorf("run sql totalTopUpAmountQuery: %w", err)
	}

	defer func() {
		if closeErr := rows.Close(); err != nil {
			err = errors.Join(err, fmt.Errorf("rows close err: %w", closeErr))
		}
	}()

	for rows.Next() {
		var currency string
		var amount decimal.Decimal
		if err := rows.Scan(&currency, &amount); err != nil {
			return &models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return &models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		(*totalTopupAmount)[curEnum] = amount
	}
	return totalTopupAmount, nil
}
