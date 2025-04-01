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

func (c *TransactionService) IsEnoughMoney(amount, balance int) (bool, error) {
	return c.repo.IsEnoughMoney(amount, balance)
}

func (c *TransactionService) SendMoney(username string, amount int, recipient string) error {
	return c.repo.SendMoney(username, amount, recipient)
}
func (c *TransactionService) TakeBalance(username string) (int, error) {
	return c.repo.TakeBalance(username)
}
