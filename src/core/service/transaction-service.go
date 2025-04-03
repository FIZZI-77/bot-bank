package service

import (
	"fmt"
	"tg_transaction/src/core/repository"
)

type TransactionService struct {
	repo       repository.TransactionRepo
	actionRepo repository.UserActionsRepo
}

func NewTransactionService(repo repository.TransactionRepo, actionRepo repository.UserActionsRepo) *TransactionService {
	return &TransactionService{repo: repo, actionRepo: actionRepo}
}

func (c *TransactionService) IsEnoughMoney(amount, balance float64) bool {
	if amount > balance {
		return false
	}
	return true
}

func (c *TransactionService) SendMoney(username string, amount float64, recipient string) error {
	recipientTgID, err := c.actionRepo.TakeUserTgID(recipient)
	userTgID, err := c.actionRepo.TakeUserTgID(username)
	if err != nil {
		return fmt.Errorf("take user tgID %s failed: %v", username, err)
	}

	balance, err := c.TakeBalance(username)

	if err != nil {
		return fmt.Errorf("service: SendMoney : error taking balance: %v", err)
	}

	if !c.IsEnoughMoney(amount, balance) {
		return fmt.Errorf("you have enough money on your balance")
	}
	return c.repo.SendMoney(userTgID, amount, recipientTgID)
}
func (c *TransactionService) TakeBalance(username string) (float64, error) {
	senderID, err := c.actionRepo.TakeUserTgID(username)
	if err != nil {
		return 0, fmt.Errorf("take user tgID %s failed: %v", username, err)
	}
	return c.repo.TakeBalance(senderID)
}
