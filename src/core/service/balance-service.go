package service

import "fmt"

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

func (b *BalanceService) TakeTotalBalance(username string) (float64, error) {
	var totalBalance float64

	totalTopUp, err := b.top.GetTotalTopupAmount(username)
	if err != nil {
		return totalBalance, fmt.Errorf("Balance-Service:get total topup amount failed: %v", err)
	}
	totalTransactionAmount, err := b.tr.GetTotalTransactionAmount(username)
	if err != nil {
		return totalBalance, fmt.Errorf("Balance-Service:get total transaction amount failed: %v", err)
	}
	totalBalance = totalTopUp + totalTransactionAmount
	return totalBalance, nil
}
