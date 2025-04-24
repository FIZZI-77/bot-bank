package service

import (
	"context"
	"github.com/mymmrac/telego"
	"github.com/shopspring/decimal"
	"tgtransaction/src/core/models"
	"tgtransaction/src/core/repository"
)

type TopUpMoney interface {
	PersistTopUp(ctx context.Context, username string, amount decimal.Decimal, cur models.CurrencyEnum) error
	GetTotalTopupAmount(ctx context.Context, username string) (*models.Balance, error)
}

type Transaction interface {
	PersistTransaction(ctx context.Context, username string, amount decimal.Decimal, recipient string, cur models.CurrencyEnum) error
	IsEnoughMoney(amount decimal.Decimal, balance models.Balance, cur models.CurrencyEnum) bool
	GetTotalTransactionAmount(ctx context.Context, username string) (*models.Balance, error)
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type UserActions interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username string, tgID int64) error
}

type Balance interface {
	TakeTotalBalance(ctx context.Context, username string) (*models.Balance, error)
}
type Service struct {
	TopUpMoney
	Transaction
	BotActions
	UserActions
	Balance
}

func NewService(repos *repository.Repository) *Service {

	transactionService := NewTransactionService(
		repos.TransactionRepo,
		repos.UserActionsRepo,
		nil,
	)

	topUpService := NewTopUpService(
		repos.TopUpMoneyRepo,
		repos.UserActionsRepo,
	)

	balanceService := NewBalanceService(
		topUpService,
		transactionService,
	)

	transactionService.SetBalance(balanceService)
	return &Service{
		topUpService,
		transactionService,
		NewBotActionsService(),
		NewUserActionService(repos.UserActionsRepo),
		balanceService,
	}
}
