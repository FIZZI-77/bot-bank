package repository

import (
	"database/sql"
	"fmt"
)

type TopUpPostgres struct {
	db *sql.DB
}

func NewTopUpPostgres(db *sql.DB) *TopUpPostgres {
	return &TopUpPostgres{db: db}
}

func (c *TopUpPostgres) PersistTopUp(userId int64, amount float64) error {

	const topUpQuery = "INSERT INTO topup_history (user_telegram_id,amount) VALUES ($1,$2)"
	_, err := c.db.Exec(topUpQuery, userId, amount)
	if err != nil {
		return fmt.Errorf("top-up-repo: TopUpMoney() : topUp money failed: %v", err)
	}
	return nil
}

func (c *TopUpPostgres) GetTotalTopupAmount(userTgID int64) (float64, error) {
	var totalTopupAmount float64

	const totalTopUpAmountQuery = `SELECT
		COALESCE(SUM(CASE WHEN user_telegram_id = $1 THEN amount ELSE 0 END),0)
		AS totalTopupAmount
	FROM topup_history`

	err := c.db.QueryRow(totalTopUpAmountQuery, userTgID).Scan(&totalTopupAmount)
	if err != nil {
		return 0, fmt.Errorf("run sql totalTopUpAmountQuery: %w", err)
	}
	return totalTopupAmount, nil
}
