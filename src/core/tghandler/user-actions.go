package tghandler

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (h *Handler) addUser(ctx context.Context, username string, tgID int64) {

	isExist, err := h.service.UserActions.UserExistsByUsername(ctx, username)

	if err != nil {
		logrus.Errorf("user-actions handler: addUser() :error check user exists : %s\n", err.Error())
	}

	if !isExist {
		err = h.service.UserActions.PersistUser(ctx, username, tgID)
		if err != nil {
			logrus.Errorf("user-actions handler: addUser() : add user failed: %v\n", err)
		}
	} else {
		logrus.Errorf("user-actions handler: addUser() : add user failed: already exists")
	}

}
