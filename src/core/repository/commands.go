package repository

import (
	"database/sql"
	"github.com/mymmrac/telego"
)

type CommandPostgres struct {
	db *sql.DB
}

func NewCommandPostgres(db *sql.DB) *CommandPostgres {
	return &CommandPostgres{db: db}
}

func (c *CommandPostgres) SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error {
	return nil
}

func (c *CommandPostgres) TopUpMoney(bot *telego.Bot, chatID int64, amount string) error {
	return nil
}
