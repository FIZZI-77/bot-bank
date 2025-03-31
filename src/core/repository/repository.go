package repository

import (
	"database/sql"
	"github.com/mymmrac/telego"
)

type TopUpMoney interface {
	TopUpMoney(bot *telego.Bot, chatID int64, amount string) error
}

type Transaction interface {
	SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error
}

type Repository struct {
	TopUpMoney
	Transaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewTopUpPostgres(db),
		NewTransactionPostgres(db),
	}
}
