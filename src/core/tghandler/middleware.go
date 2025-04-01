package tghandler

import (
	"github.com/sirupsen/logrus"
)

// Проверка что сумма положительна

func (h *Handler) checkCommand(command string, parts []string, amount int) bool {

	if amount <= 0 {
		logrus.Error("ведена неверная сумма для перевода")
		return false
	}
	switch command {
	case "/send":
		if len(parts) != 3 {
			logrus.Error("неправильное кол-во аргументов в комманде /send")
			return false
		}

	case "/top-up":
		if len(parts) != 2 {
			logrus.Error("неправильное кол-во аргументов в комманде /top-up")
			return false
		}

	}
	return true

}
