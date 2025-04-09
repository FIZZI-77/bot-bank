package service

import (
	"fmt"
	"tg_transaction/src/core/repository"
)

type TopUpService struct {
	repo       repository.TopUpMoneyRepo
	actionRepo repository.UserActionsRepo
}

func NewTopUpService(repo repository.TopUpMoneyRepo, actionRepo repository.UserActionsRepo) *TopUpService {
	return &TopUpService{repo: repo, actionRepo: actionRepo}
}

func (c *TopUpService) TopUpMoney(username string, amount float64) error {

	senderID, err := c.actionRepo.TakeUserTgID(username)
	if err != nil {
		return fmt.Errorf("top-up-service : TopUpMoney() : take user tgID %s failed: %v", username, err)
	}

	return c.repo.TopUpMoney(senderID, amount)
}

func (c *TopUpService) GetTotalTopupAmount(username string) (float64, error) {
	senderID, err := c.actionRepo.TakeUserTgID(username)
	if err != nil {
		return 0, fmt.Errorf("top-up-service : GetTotalTopupAmount() : take user tgID %s failed: %v", username, err)
	}
	return c.repo.GetTotalTopupAmount(senderID)
}
