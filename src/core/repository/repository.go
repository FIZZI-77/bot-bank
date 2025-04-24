package repository

import (
	"context"
	"database/sql"
	"tgtransaction/src/core/models"
)

type TopUpMoneyRepo interface {
	PersistTopUp(model *models.TopUpModel) error
	GetTotalTopupAmount(ctx context.Context, userTgID int64) (*models.Balance, error)
}

type TransactionRepo interface {
	PersistTransaction(model *models.TransactionModel) error
	GetTotalSentAmount(ctx context.Context, userTgID string) (*models.Balance, error)
	GetTotalReceivedAmount(ctx context.Context, userTgID int64) (*models.Balance, error)
}

type UserActionsRepo interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username, userid string, tgID int64) error
	GetUserTgIDByUsername(ctx context.Context, username string) (int64, error)
	GetIDByUsername(ctx context.Context, username string) (string, error)
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
