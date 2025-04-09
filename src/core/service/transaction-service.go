package service

import (
	"fmt"
	"tg_transaction/src/core/repository"
)

type TransactionService struct {
	repo       repository.TransactionRepo
	actionRepo repository.UserActionsRepo
	balance    Balance
}

func NewTransactionService(repo repository.TransactionRepo, actionRepo repository.UserActionsRepo, balance Balance) *TransactionService {
	return &TransactionService{repo: repo, actionRepo: actionRepo, balance: balance}
}

func (c *TransactionService) SetBalance(balance Balance) {
	c.balance = balance
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
		return fmt.Errorf("transaction-service : SendMoney() :take user tgID %s failed: %v", username, err)
	}

	balance, err := c.balance.TakeTotalBalance(username)

	if err != nil {
		return fmt.Errorf("transaction-service: SendMoney : error taking balance: %v", err)
	}

	if !c.IsEnoughMoney(amount, balance) {
		return fmt.Errorf("transaction-service : SendMoney() : you have not enough money on your balance")

	}
	return c.repo.SendMoney(userTgID, amount, recipientTgID)
}

func (c *TransactionService) GetTotalTransactionAmount(username string) (float64, error) {
	var totalAmount float64
	senderID, err := c.actionRepo.TakeUserTgID(username)
	if err != nil {
		return 0, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : take user tgID %s failed: %v", username, err)
	}

	totalReceivedAmount, err := c.repo.GetTotalReceivedAmount(senderID)
	if err != nil {
		return 0, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalReceivedAmount failed: %v", err)
	}
	totalSentAmount, err := c.repo.GetTotalSentAmount(senderID)
	if err != nil {
		return 0, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalSentAmount failed: %v", err)
	}

	totalAmount = totalReceivedAmount - totalSentAmount

	return totalAmount, nil
}
