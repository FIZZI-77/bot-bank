package repository

import (
	"database/sql"
	"github.com/mymmrac/telego"
)

type TopUpPostgres struct {
	db *sql.DB
}

func NewTopUpPostgres(db *sql.DB) *TopUpPostgres {
	return &TopUpPostgres{db: db}
}

func (c *TopUpPostgres) TopUpMoney(bot *telego.Bot, chatID int64, amount string) error {
	return nil
}
