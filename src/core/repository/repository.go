package repository

import (
	"database/sql"
)

type TopUpMoneyRepo interface {
	PersistTopUp(userId int64, amount float64) error
	GetTotalTopupAmount(userTgID int64) (float64, error)
}

type TransactionRepo interface {
	PersistTransaction(username int64, amount float64, recipient int64) error
	GetTotalSentAmount(userTgID int64) (float64, error)
	GetTotalReceivedAmount(userTgID int64) (float64, error)
}

type UserActionsRepo interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username, userid string, tgID int64) error
	GetUserTgIDByUsername(username string) (int64, error)
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
