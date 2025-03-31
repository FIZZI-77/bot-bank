package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TopUpService struct {
	repo repository.TopUpMoney
}

func NewTopUpService(repo repository.TopUpMoney) *TopUpService {
	return &TopUpService{repo: repo}
}

func (c *TopUpService) TopUpMoney(bot *telego.Bot, chatID int64, amount string) error {
	return c.repo.TopUpMoney(bot, chatID, amount)
}
