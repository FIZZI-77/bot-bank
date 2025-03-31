package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type Commands interface {
	SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error
	TopUpMoney(bot *telego.Bot, chatID int64, amount string) error
	SendMessage(bot *telego.Bot, chatID int64, msg string) error
}

type Service struct {
	Commands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewCommandService(repos.Commands),
	}
}
