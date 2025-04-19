package service

import (
	"fmt"
	"github.com/google/uuid"
	"tg_transaction/src/core/models"
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

func (c *TransactionService) IsEnoughMoney(amount float64, balance models.Balance, cur models.CurrencyEnum) bool {
	if amount > balance[cur] {
		return false
	}
	return true
}

func (c *TransactionService) PersistTransaction(username string, amount float64, recipient string, cur models.CurrencyEnum) (int, error) {
	transactionId := uuid.New().String()
	recipientTgID, err := c.actionRepo.GetUserTgIDByUsername(recipient)
	userId, err := c.actionRepo.GetUUIDByUsername(username)
	if err != nil {
		return 1, fmt.Errorf("transaction-service : PersistTransaction() :take user uuid %s failed: %v", username, err)
	}

	balance, err := c.balance.TakeTotalBalance(username)

	if err != nil {
		return 1, fmt.Errorf("transaction-service: PersistTransaction : error taking balance: %v", err)
	}

	if !c.IsEnoughMoney(amount, balance, cur) {
		return 2, fmt.Errorf("transaction-service : PersistTransaction() : you have not enough money on your balance")

	}
	return 0, c.repo.PersistTransaction(userId, amount, recipientTgID, transactionId, cur)
}

func (c *TransactionService) GetTotalTransactionAmount(username string) (models.Balance, error) {
	totalAmount := make(models.Balance)
	senderID, err := c.actionRepo.GetUUIDByUsername(username)
	if err != nil {
		return models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : take user uuid %s failed: %v", username, err)
	}

	recipientID, err := c.actionRepo.GetUserTgIDByUsername(username)
	if err != nil {
		return models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : take user tgID %s failed: %v", username, err)
	}
	totalReceivedAmount, err := c.repo.GetTotalReceivedAmount(recipientID)
	if err != nil {
		return models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalReceivedAmount failed: %v", err)
	}
	totalSentAmount, err := c.repo.GetTotalSentAmount(senderID)
	if err != nil {
		return models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalSentAmount failed: %v", err)
	}

	for currency, amount := range totalReceivedAmount {
		totalAmount[currency] = amount
	}

	for currency, amount := range totalSentAmount {
		totalAmount[currency] -= amount
	}
	return totalAmount, nil
}
