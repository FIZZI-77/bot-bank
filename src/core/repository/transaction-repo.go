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

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	operationType := "send"

	const sendQuery = "INSERT INTO operations_history (sender,operation_type,amount,recipient) VALUES ($1,$2,$3,$4)"

	_, err = tx.Exec(sendQuery, username, operationType, amount, recipient)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("transaction-repo: SendMoney() : transaction rollback failed: %v", err)
		}
		return fmt.Errorf("transaction-repo: SendMoney() : send money failed: %v", err)
	}
	return tx.Commit()
}

func (c *TransactionPostgres) TakeBalance(userTgID int64) (float64, error) {

	var balance float64

	const takeSendOperationsQuery = `SELECT 
		COALESCE(SUM(CASE WHEN operation_type = 'top-up' THEN amount ELSE 0 END),0) +
		COALESCE(SUM(CASE WHEN operation_type = 'send' and recipient = $1 THEN amount ELSE 0 END ),0) -
		COALESCE(SUM(CASE WHEN operation_type = 'send' and sender = $1 THEN amount ELSE 0 END ),0)
		AS balance
	FROM operations_history
	`
	err := c.db.QueryRow(takeSendOperationsQuery, userTgID).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("transaction-repo: TakeBalance() : TakeBalance failed: %v", err)
	}
	return balance, nil
}
