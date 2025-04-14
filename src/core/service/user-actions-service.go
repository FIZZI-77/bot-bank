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

func (c *UserActionService) UserExistsByUsername(username string) (bool, error) {
	return c.repo.UserExistsByUsername(username)
}

func (c *UserActionService) PersistUser(username string, tgID int64) error {
	userid := uuid.New().String()
	return c.repo.PersistUser(username, userid, tgID)
}
