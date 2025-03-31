package repository

import (
	"database/sql"
	"github.com/mymmrac/telego"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (c *TransactionPostgres) SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error {
	return nil
}
