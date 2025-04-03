package service

import (
	"github.com/google/uuid"
	"tg_transaction/src/core/repository"
)

type UserActionService struct {
	repo repository.UserActionsRepo
}

func NewUserActionService(repo repository.UserActionsRepo) *UserActionService {
	return &UserActionService{repo: repo}
}

func (c *UserActionService) IsUserExists(username string) (bool, error) {
	return c.repo.IsUserExists(username)
}

func (c *UserActionService) AddUser(username string, tgID int64) error {
	userid := uuid.New().String()
	return c.repo.AddUser(username, userid, tgID)
}
