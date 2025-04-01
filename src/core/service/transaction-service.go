package service

import (
	"tg_transaction/src/core/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (c *TransactionService) SendMoney(username string, amount int, recipient string) error {
	return c.repo.SendMoney(username, amount, recipient)
}
