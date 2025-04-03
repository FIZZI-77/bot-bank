package repository

import (
	"database/sql"
)

type TopUpMoney interface {
	TopUpMoney(username string, amount int) error
}

type Transaction interface {
	SendMoney(username string, amount int, recipient string) error
	TakeBalance(username string) (int, error)
}

type UserActions interface {
	IsUserExists(username string) (bool, error)
	AddUser(username string, tgID int) error
}

type Repository struct {
	TopUpMoney
	Transaction
	UserActions
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewTopUpPostgres(db),
		NewTransactionPostgres(db),
		NewUserActionPostgres(db),
	}
}
