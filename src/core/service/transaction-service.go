package service

import (
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (c *TransactionService) SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error {
	return c.repo.SendMoney(bot, chatID, amount, recipient)
}
