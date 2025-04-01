package tghandler

import (
	"github.com/sirupsen/logrus"
)

func (h *Handler) addUser(username string) {

	isExist, err := h.service.UserActions.IsUserExists(username)

	if err != nil {
		logrus.Errorf("Ошибка проверки : %s\n", err.Error())
	}

	if isExist {
		err = h.service.UserActions.AddUser(username)
		if err != nil {
			logrus.Errorf("add user failed: %v\n", err)
		}
	} else {
		logrus.Errorf("add user failed: already exists")
	}

}
