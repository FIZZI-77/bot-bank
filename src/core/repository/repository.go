package repository

import (
	"database/sql"
	"tg_transaction/src/core/models"
)

type TopUpMoneyRepo interface {
	PersistTopUp(userId int64, amount float64, topUpId string, cur models.CurrencyEnum) error
	GetTotalTopupAmount(userTgID int64) (models.Balance, error)
}

type TransactionRepo interface {
	PersistTransaction(userId string, amount float64, recipient int64, transactionId string, cur models.CurrencyEnum) error
	GetTotalSentAmount(userTgID string) (models.Balance, error)
	GetTotalReceivedAmount(userTgID int64) (models.Balance, error)
}

type UserActionsRepo interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username, userid string, tgID int64) error
	GetUserTgIDByUsername(username string) (int64, error)
	GetUUIDByUsername(username string) (string, error)
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
