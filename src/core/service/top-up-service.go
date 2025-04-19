package service

import (
	"fmt"
	"github.com/google/uuid"
	"tg_transaction/src/core/models"
	"tg_transaction/src/core/repository"
)

type TopUpService struct {
	repo       repository.TopUpMoneyRepo
	actionRepo repository.UserActionsRepo
}

func NewTopUpService(repo repository.TopUpMoneyRepo, actionRepo repository.UserActionsRepo) *TopUpService {
	return &TopUpService{repo: repo, actionRepo: actionRepo}
}

func (c *TopUpService) PersistTopUp(username string, amount float64, cur models.CurrencyEnum) error {
	topUpId := uuid.New().String()
	senderID, err := c.actionRepo.GetUserTgIDByUsername(username)
	if err != nil {
		return fmt.Errorf("top-up-service : TopUpMoney() : take user tgID %s failed: %v", username, err)
	}

	return c.repo.PersistTopUp(senderID, amount, topUpId, cur)
}

func (c *TopUpService) GetTotalTopupAmount(username string) (models.Balance, error) {
	senderID, err := c.actionRepo.GetUserTgIDByUsername(username)
	if err != nil {
		return models.Balance{}, fmt.Errorf("top-up-service : GetTotalTopupAmount() : take user tgID %s failed: %v", username, err)
	}
	return c.repo.GetTotalTopupAmount(senderID)
}
