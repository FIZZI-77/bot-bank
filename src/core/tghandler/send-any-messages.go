package tghandler

import (
	"fmt"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"tg_transaction/src/core/tghandler/messages"
)

func (h *Handler) sendStartMessage(bot *telego.Bot, chatID int64) error {
	err := h.service.SendMessage(bot, chatID, messages.MsgHello)
	if err != nil {
		return fmt.Errorf("handler sendStartMessage: cant't send start message %v", err)
	}
	return nil
}

func (h *Handler) sendHelpMessage(bot *telego.Bot, chatID int64) error {
	err := h.service.SendMessage(bot, chatID, messages.MsgHelp)
	if err != nil {
		return fmt.Errorf("handler sendHelpMessage: cant't send help message %v", err)
	}
	return nil
}

func (h *Handler) sendUnknownMessage(bot *telego.Bot, chatID int64) {
	err := h.service.SendMessage(bot, chatID, messages.MsgUnknownCommand)
	if err != nil {
		logrus.Errorf("handler sendUnknownMessage: cant't send unknown message %v", err)
	}
}

func (h *Handler) takeBalanceMessage(bot *telego.Bot, chatID int64, username string) {
	balance := h.takeBalance(username)
	err := h.service.SendMessage(bot, chatID, messages.MsgBalance+fmt.Sprintf("%d", balance))
	if err != nil {
		logrus.Errorf("handler takeBalanceMessage: cant't send balance message %v", err)
	}
}
