package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"tgtransaction/src/core/models"
)

type TopUpMoneyRepo interface {
	PersistTopUp(ctx context.Context, model models.TopUpModel) error
	GetTotalTopupAmount(ctx context.Context, userTgID int64) (models.Balance, error)
}

type TransactionRepo interface {
	PersistTransaction(ctx context.Context, model models.TransactionModel) error
	GetTotalSentAmount(ctx context.Context, userTgID uuid.UUID) (models.Balance, error)
	GetTotalReceivedAmount(ctx context.Context, userTgID int64) (models.Balance, error)
}

type UserActionsRepo interface {
	UserExistsByUsername(ctx context.Context, username string) (bool, error)
	PersistUser(ctx context.Context, username string, userid uuid.UUID, tgID int64) error
	GetUserTgIDByUsername(ctx context.Context, username string) (int64, error)
	GetIDByUsername(ctx context.Context, username string) (uuid.UUID, error)
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
