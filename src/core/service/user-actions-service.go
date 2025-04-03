package service

import "tg_transaction/src/core/repository"

type UserActionService struct {
	repo repository.UserActions
}

func NewUserActionService(repo repository.UserActions) *UserActionService {
	return &UserActionService{repo: repo}
}

func (c *UserActionService) IsUserExists(username string) (bool, error) {
	return c.repo.IsUserExists(username)
}

func (c *UserActionService) AddUser(username string, tgID int) error {
	return c.repo.AddUser(username, tgID)
}
