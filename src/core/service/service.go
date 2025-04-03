package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TopUpMoney interface {
	TopUpMoney(username string, amount float64) error
}

type Transaction interface {
	SendMoney(username string, amount float64, recipient string) error
	TakeBalance(username string) (float64, error)
	IsEnoughMoney(amount, balance float64) bool
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type UserActions interface {
	IsUserExists(username string) (bool, error)
	AddUser(username string, tgID int64) error
}
type Service struct {
	TopUpMoney
	Transaction
	BotActions
	UserActions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewTopUpService(repos.TopUpMoneyRepo, repos.UserActionsRepo),
		NewTransactionService(repos.TransactionRepo, repos.UserActionsRepo),
		NewBotActionsService(),
		NewUserActionService(repos.UserActionsRepo),
	}
}
