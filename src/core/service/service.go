package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TopUpMoney interface {
	TopUpMoney(username string, amount int) error
}

type Transaction interface {
	SendMoney(username string, amount int, recipient string) error
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type UserActions interface {
	IsUserExists(username string) (bool, error)
	AddUser(username string) error
}
type Service struct {
	TopUpMoney
	Transaction
	BotActions
	UserActions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewTopUpService(repos.TopUpMoney),
		NewTransactionService(repos.Transaction),
		NewBotActionsService(),
		NewUserActionService(repos.UserActions),
	}
}
