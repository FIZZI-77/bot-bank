package repository

import (
	"database/sql"
	"github.com/mymmrac/telego"
)

type Commands interface {
	SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error
	TopUpMoney(bot *telego.Bot, chatID int64, amount string) error
}

type Repository struct {
	Commands
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewCommandPostgres(db),
	}
}
