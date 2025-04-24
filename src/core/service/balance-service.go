package service

import (
	"context"
	"fmt"
	"tgtransaction/src/core/models"
)

type BalanceService struct {
	top TopUpMoney
	tr  Transaction
}

func NewBalanceService(top TopUpMoney, tr Transaction) *BalanceService {
	return &BalanceService{
		top: top,
		tr:  tr,
	}
}

func (b *BalanceService) TakeTotalBalance(ctx context.Context, username string) (*models.Balance, error) {
	totalBalance := &models.Balance{}

	totalTopUp, err := b.top.GetTotalTopupAmount(ctx, username)
	if err != nil {
		return totalBalance, fmt.Errorf("Balance-Service:get total topup amount failed: %v", err)
	}
	totalTransactionAmount, err := b.tr.GetTotalTransactionAmount(ctx, username)
	if err != nil {
		return totalBalance, fmt.Errorf("Balance-Service:get total transaction amount failed: %v", err)
	}
	for currency, amount := range *totalTopUp {
		(*totalBalance)[currency] = amount
	}

	for currency, amount := range *totalTransactionAmount {
		(*totalBalance)[currency] = (*totalBalance)[currency].Add(amount)
	}
	return totalBalance, nil
}
