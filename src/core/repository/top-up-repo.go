package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"tg_transaction/src/core/models"
)

type TopUpPostgres struct {
	db *sql.DB
}

func NewTopUpPostgres(db *sql.DB) *TopUpPostgres {
	return &TopUpPostgres{db: db}
}

func (c *TopUpPostgres) PersistTopUp(userId int64, amount float64, topUpId string, cur models.CurrencyEnum) error {
	var currency string
	currency, err := FromCurrencyEnum(cur)
	if err != nil {
		return err
	}

	const topUpQuery = "INSERT INTO topup (id,user_telegram_id,amount,currency) VALUES ($1,$2,$3,$4)"
	_, err = c.db.Exec(topUpQuery, topUpId, userId, amount, currency)
	if err != nil {
		return fmt.Errorf("run sql topUpQuery: %w", err)
	}
	return nil
}

func (c *TopUpPostgres) GetTotalTopupAmount(userTgID int64) (models.Balance, error) {
	totalTopupAmount := make(models.Balance)
	
	const totalTopUpAmountQuery = `SELECT
    currency,
	COALESCE(SUM(CASE WHEN user_telegram_id = $1 THEN amount ELSE 0 END),0)
	FROM topup
	GROUP BY currency`

	rows, err := c.db.Query(totalTopUpAmountQuery, userTgID)
	if err != nil {
		return models.Balance{}, fmt.Errorf("run sql totalTopUpAmountQuery: %w", err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			logrus.Error("Error rows.Close(): %v", err)
		}
	}()

	for rows.Next() {
		var currency string
		var amount float64
		if err := rows.Scan(&currency, &amount); err != nil {
			return models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		totalTopupAmount[curEnum] = amount
	}
	return totalTopupAmount, nil
}
