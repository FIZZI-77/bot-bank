package repository

import (
	"database/sql"
)

type TopUpMoneyRepo interface {
	TopUpMoney(userId int64, amount float64) error
	GetTotalTopupAmount(userTgID int64) (float64, error)
}

type TransactionRepo interface {
	SendMoney(username int64, amount float64, recipient int64) error
	GetTotalSentAmount(userTgID int64) (float64, error)
	GetTotalReceivedAmount(userTgID int64) (float64, error)
}

type UserActionsRepo interface {
	IsUserExists(username string) (bool, error)
	AddUser(username, userid string, tgID int64) error
	TakeUserTgID(username string) (int64, error)
}

type Repository struct {
	UserActionsRepo
	TopUpMoneyRepo
	TransactionRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewUserActionPostgres(db),
		NewTopUpPostgres(db),
		NewTransactionPostgres(db),
	}
}
