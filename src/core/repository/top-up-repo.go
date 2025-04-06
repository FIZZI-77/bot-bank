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

func (c *TopUpPostgres) TopUpMoney(userId int64, amount float64) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	operation := "top-up"
	const topUpQuery = "INSERT INTO operations_history (sender,operation_type,amount) VALUES ($1,$2,$3)"
	_, err = tx.Exec(topUpQuery, userId, operation, amount)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("top-up-repo: TopUpMoney() : topUp rollback failed: %v", err)
		}
		return fmt.Errorf("top-up-repo: TopUpMoney() : topUp money failed: %v", err)
	}
	return tx.Commit()
}
