package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TopUpMoney interface {
	PersistTopUp(username string, amount float64) error
	GetTotalTopupAmount(username string) (float64, error)
}

type Transaction interface {
	PersistTransaction(username string, amount float64, recipient string) error
	IsEnoughMoney(amount, balance float64) bool
	GetTotalTransactionAmount(username string) (float64, error)
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type UserActions interface {
	UserExistsByUsername(username string) (bool, error)
	PersistUser(username string, tgID int64) error
}

type Balance interface {
	TakeTotalBalance(username string) (float64, error)
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
		nil, // временно nil
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
