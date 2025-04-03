package service

import (
	"fmt"
	"tg_transaction/src/core/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (c *TransactionService) IsEnoughMoney(amount, balance int) bool {
	if amount > balance {
		return false
	}
	return true
}

func (c *TransactionService) SendMoney(username string, amount int, recipient string) error {

	balance, err := c.TakeBalance(username)

	if err != nil {
		return fmt.Errorf("service: SendMoney : error taking balance: %v", err)
	}

	if c.IsEnoughMoney(amount, balance) {
		return fmt.Errorf("you have enough money on your balance")
	}
	return c.repo.SendMoney(username, amount, recipient)
}
func (c *TransactionService) TakeBalance(username string) (int, error) {

	return c.repo.TakeBalance(username)
}
