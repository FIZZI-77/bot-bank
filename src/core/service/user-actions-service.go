package service

import (
	"context"
	"github.com/google/uuid"
	"tgtransaction/src/core/repository"
)

type UserActionService struct {
	repo repository.UserActionsRepo
}

func NewUserActionService(repo repository.UserActionsRepo) *UserActionService {
	return &UserActionService{repo: repo}
}

func (c *UserActionService) UserExistsByUsername(ctx context.Context, username string) (bool, error) {
	return c.repo.UserExistsByUsername(ctx, username)
}

func (c *UserActionService) PersistUser(ctx context.Context, username string, tgID int64) error {
	userid := uuid.New()
	return c.repo.PersistUser(ctx, username, userid, tgID)
}
