package repository

import (
	"database/sql"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (c *TransactionPostgres) SendMoney(username string, amount int, recipient string) error {
	return nil
}
