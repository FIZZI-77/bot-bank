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

func (c *TopUpPostgres) TopUpMoney(username string, amount int) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	operation := "top-up"
	topUpQuery := fmt.Sprintf("INSERT INTO operations (operation_author,operation,amount) VALUES ($1,$2,$3)")
	_, err = tx.Exec(topUpQuery, username, operation, amount)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("topUp rollback failed: %v", err)
		}
		return fmt.Errorf("topUp money failed: %v", err)
	}
	return tx.Commit()
}
