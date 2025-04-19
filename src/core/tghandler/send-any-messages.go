package tghandler

import (
	"fmt"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"strings"
	"tg_transaction/src/core/tghandler/messages"
)

func (h *Handler) sendStartMessage(bot *telego.Bot, chatID int64) error {
	err := h.service.SendMessage(bot, chatID, messages.MsgHello)
	if err != nil {
		return fmt.Errorf("send-any-messages handler : sendStartMessage() : cant't send start message %v", err)
	}
	return nil
}

func (h *Handler) sendHelpMessage(bot *telego.Bot, chatID int64) error {
	err := h.service.SendMessage(bot, chatID, messages.MsgHelp)
	if err != nil {
		return fmt.Errorf("send-any-messages handler :sendHelpMessage(): cant't send help message %v", err)
	}
	return nil
}

func (h *Handler) sendUnknownMessage(bot *telego.Bot, chatID int64) {
	err := h.service.SendMessage(bot, chatID, messages.MsgUnknownCommand)
	if err != nil {
		logrus.Errorf("send-any-messages handler : sendUnknownMessage(): cant't send unknown message %v", err)
	}
}

func (h *Handler) takeBalanceMessage(bot *telego.Bot, chatID int64, username string) {
	balance := h.takeBalance(username)
	var parts []string
	for currency, amount := range balance {
		part := fmt.Sprintf("%s: %.2f", currencyToString[currency], amount)
		parts = append(parts, part)
	}
	balanceStr := strings.Join(parts, "\n")
	err := h.service.SendMessage(bot, chatID, messages.MsgBalance+"\n"+balanceStr)
	if err != nil {
		logrus.Errorf("send-any-messages handler : takeBalanceMessage(): cant't send balance message %v", err)
	}
}

func (h *Handler) sendErrorCurrencyMessage(bot *telego.Bot, chatID int64) {
	err := h.service.SendMessage(bot, chatID, messages.MsgErrCurrency)
	if err != nil {
		logrus.Errorf("send-any-messages handler : sendErrorCurrencyMessage(): cant't send error currancy message %v", err)
	}
}
