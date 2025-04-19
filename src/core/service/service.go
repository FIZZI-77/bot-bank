package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/models"
	"tg_transaction/src/core/repository"
)

type TopUpMoney interface {
	PersistTopUp(username string, amount float64, cur models.CurrencyEnum) error
	GetTotalTopupAmount(username string) (models.Balance, error)
}

type Transaction interface {
	PersistTransaction(username string, amount float64, recipient string, cur models.CurrencyEnum) (int, error)
	IsEnoughMoney(amount float64, balance models.Balance, cur models.CurrencyEnum) bool
	GetTotalTransactionAmount(username string) (models.Balance, error)
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type UserActions interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username string, tgID int64) error
}

type Balance interface {
	TakeTotalBalance(username string) (models.Balance, error)
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
