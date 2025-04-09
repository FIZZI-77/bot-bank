package repository

import (
	"database/sql"
	"fmt"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (c *TransactionPostgres) SendMoney(username int64, amount float64, recipient int64) error {

	const sendQuery = "INSERT INTO transactions_history (sender,amount,recipient) VALUES ($1,$2,$3)"

	_, err := c.db.Exec(sendQuery, username, amount, recipient)
	if err != nil {

		return fmt.Errorf("run sql sendQuery: %w", err)
	}
	return nil
}

func (c *TransactionPostgres) GetTotalSentAmount(userTgID int64) (float64, error) {

	var totalSent float64

	const getTotalSendQuery = `SELECT
		COALESCE(SUM(CASE WHEN sender = $1 THEN amount ELSE 0 END ),0)
		AS totalSent
	FROM transactions_history
`
	err := c.db.QueryRow(getTotalSendQuery, userTgID).Scan(&totalSent)
	if err != nil {
		return 0, fmt.Errorf("run sql getTotalSendQuery: %w", err)
	}
	return totalSent, nil
}

func (c *TransactionPostgres) GetTotalReceivedAmount(userTgID int64) (float64, error) {
	var totalReceived float64

	const getTotalReceivedQuery = `SELECT
		COALESCE(SUM(CASE WHEN recipient = $1 THEN amount ELSE 0 END ),0)
		AS totalReceived
	FROM transactions_history
`
	err := c.db.QueryRow(getTotalReceivedQuery, userTgID).Scan(&totalReceived)
	if err != nil {
		return 0, fmt.Errorf("run sql getTotalReceivedQuery: %w", err)
	}
	return totalReceived, nil
}
