package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"tgtransaction/src/core/models"
	"tgtransaction/src/core/repository"
)

type TopUpService struct {
	repo       repository.TopUpMoneyRepo
	actionRepo repository.UserActionsRepo
}

func NewTopUpService(repo repository.TopUpMoneyRepo, actionRepo repository.UserActionsRepo) *TopUpService {
	return &TopUpService{repo: repo, actionRepo: actionRepo}
}

func (c *TopUpService) PersistTopUp(ctx context.Context, username string, amount decimal.Decimal, cur models.CurrencyEnum) error {
	topUpId := uuid.New()
	senderID, err := c.actionRepo.GetUserTgIDByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("top-up-service : TopUpMoney() : take user tgID %s failed: %v", username, err)
	}
	topUp := &models.TopUpModel{
		UserId:  senderID,
		TopUpId: topUpId,
		Amount:  amount,
		Cur:     cur,
	}
	return c.repo.PersistTopUp(ctx, *topUp)
}

func (c *TopUpService) GetTotalTopupAmount(ctx context.Context, username string) (models.Balance, error) {
	senderID, err := c.actionRepo.GetUserTgIDByUsername(ctx, username)
	if err != nil {
		return models.Balance{}, fmt.Errorf("top-up-service : GetTotalTopupAmount() : take user tgID %s failed: %v", username, err)
	}
	return c.repo.GetTotalTopupAmount(ctx, senderID)
}
