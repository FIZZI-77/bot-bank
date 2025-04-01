package service

import (
	"tg_transaction/src/core/repository"
)

type TopUpService struct {
	repo repository.TopUpMoney
}

func NewTopUpService(repo repository.TopUpMoney) *TopUpService {
	return &TopUpService{repo: repo}
}

func (c *TopUpService) TopUpMoney(username string, amount int) error {
	return c.repo.TopUpMoney(username, amount)
}
