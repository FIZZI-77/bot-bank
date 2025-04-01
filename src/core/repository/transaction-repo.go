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

func (c *TransactionPostgres) IsEnoughMoney(amount, balance int) (bool, error) {
	if amount < balance {
		return false, fmt.Errorf("недостаточно средств")
	}
	return true, nil
}

func (c *TransactionPostgres) SendMoney(username string, amount int, recipient string) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	operationType := "send"
	balance, err := c.TakeBalance(username)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("transaction rollback failed: %v", err)
		}
		return fmt.Errorf("SendMoney: не удалось получить баланс: %v", err)
	}

	isEnoughMoney, err := c.IsEnoughMoney(amount, balance)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("transaction rollback failed: %v", err)
		}
		return fmt.Errorf("SendMoney: не удалось сравнить баланс: %v", err)
	}

	if !isEnoughMoney {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("transaction rollback failed: %v", err)
		}
		return fmt.Errorf("недостаточно средств")

	}

	sendQuery := fmt.Sprintf("INSERT INTO operations (operation_author,operation,amount,recipient) VALUES ($1,$2,$3,$4)")

	_, err = tx.Exec(sendQuery, username, operationType, amount, recipient)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("transaction rollback failed: %v", err)
		}
		return fmt.Errorf("send money failed: %v", err)
	}
	return tx.Commit()
}

func (c *TransactionPostgres) TakeBalance(username string) (int, error) {

	var balance int

	takeSendOperationsQuery := fmt.Sprintf(`SELECT 
	COALESCE(SUM(CASE WHEN operation = 'top-up' THEN amount ELSE 0 END),0) +
	COALESCE(SUM(CASE WHEN operation = 'send' and recipient = $1 THEN amount ELSE 0 END ),0) -
	COALESCE(SUM(CASE WHEN operation = 'send' and operation_author = $1 THEN amount ELSE 0 END ),0)
	AS balance
FROM operations
`)
	err := c.db.QueryRow(takeSendOperationsQuery, username).Scan(&balance)
	if err != nil {
		return 0, fmt.Errorf("TakeBalance failed: %v", err)
	}
	return 0, nil
}
