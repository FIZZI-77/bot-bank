package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"tgtransaction/src/core/models"
	"tgtransaction/src/core/repository"
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

func (c *TransactionService) IsEnoughMoney(amount decimal.Decimal, balance models.Balance, cur models.CurrencyEnum) bool {
	if amount.GreaterThan(balance[cur]) {
		return false
	}
	return true
}

func (c *TransactionService) PersistTransaction(ctx context.Context, username string, amount decimal.Decimal, recipient string, cur models.CurrencyEnum) error {
	transactionId := uuid.New().String()

	recipientTgID, err := c.actionRepo.GetUserTgIDByUsername(ctx, recipient)
	userId, err := c.actionRepo.GetIDByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("transaction-service : PersistTransaction() :take user uuid %s failed: %v", username, err)
	}

	balance, err := c.balance.TakeTotalBalance(ctx, username)

	if err != nil {
		return fmt.Errorf("transaction-service: PersistTransaction : error taking balance: %v", err)
	}

	if !c.IsEnoughMoney(amount, *balance, cur) {
		return &models.NotEnoughMoneyError{Message: "Not enough money"}

	}

	transaction := &models.TransactionModel{
		UserId:        userId,
		Amount:        amount,
		Recipient:     recipientTgID,
		TransactionId: transactionId,
		Cur:           cur,
	}
	return c.repo.PersistTransaction(transaction)
}

func (c *TransactionService) GetTotalTransactionAmount(ctx context.Context, username string) (*models.Balance, error) {
	totalAmount := &models.Balance{}
	senderID, err := c.actionRepo.GetIDByUsername(ctx, username)
	if err != nil {
		return &models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : take user uuid %s failed: %v", username, err)
	}

	recipientID, err := c.actionRepo.GetUserTgIDByUsername(ctx, username)
	if err != nil {
		return &models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : take user tgID %s failed: %v", username, err)
	}
	totalReceivedAmount, err := c.repo.GetTotalReceivedAmount(ctx, recipientID)
	if err != nil {
		return &models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalReceivedAmount failed: %v", err)
	}
	totalSentAmount, err := c.repo.GetTotalSentAmount(ctx, senderID)
	if err != nil {
		return &models.Balance{}, fmt.Errorf("transaction-service : GetTotalTransactionAmount() : GetTotalSentAmount failed: %v", err)
	}

	for currency, amount := range *totalReceivedAmount {
		(*totalAmount)[currency] = amount
	}

	for currency, amount := range *totalSentAmount {
		(*totalAmount)[currency] = (*totalAmount)[currency].Sub(amount)
	}
	return totalAmount, nil
}
