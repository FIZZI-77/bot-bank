package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TopUpMoney interface {
	TopUpMoney(bot *telego.Bot, chatID int64, amount string) error
}

type Transaction interface {
	SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error
}

type BotActions interface {
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type Service struct {
	TopUpMoney
	Transaction
	BotActions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewTopUpService(repos.TopUpMoney),
		NewTransactionService(repos.Transaction),
		NewBotActionsService(),
	}
}
