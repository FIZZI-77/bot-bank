package tghandler

import (
	"context"
	"errors"
	"github.com/mymmrac/telego"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"tgtransaction/src/core/models"
	"tgtransaction/src/core/tghandler/messages"
)

func (h *Handler) persistTransaction(ctx context.Context, username string, amount decimal.Decimal, recipient string, isCommandCorrect bool, bot *telego.Bot, chatID int64, cur models.CurrencyEnum) {

	if !h.isRecipientCorrect(ctx, recipient, bot, chatID) {
		return
	}
	if !isCommandCorrect {
		return
	}
	err := h.service.Transaction.PersistTransaction(ctx, username, amount, recipient, cur)

	if errors.Is(err, models.ErrNotEnoughMoney) {
		logrus.Errorf("bank-operation hendler: persistTransaction() :cant't send money: %s", err)
		if sendNotEnoughBalance := h.service.SendMessage(bot, chatID, messages.MsgNotEnoughMoney); sendNotEnoughBalance != nil {
			logrus.Errorf("bank-operation hendler: persistTransaction(): cant't send error Send  message %v", sendNotEnoughBalance)
			return
		}
		return
	}
	if err != nil {
		logrus.Errorf("bank-operation hendler: persistTransaction() :cant't send money: %s", err.Error())
		if sendErr := h.service.SendMessage(bot, chatID, messages.MsgErrorSend); sendErr != nil {
			logrus.Errorf("bank-operation hendler: persistTransaction(): cant't send error Send  message %v", sendErr)
			return

		}
		return
	}
	if err := h.service.SendMessage(bot, chatID, messages.MsgSendMoney); err != nil {
		logrus.Errorf("bank-operation hendler: persistTransaction() :cant't send Send message %v", err)
	}
}

func (h *Handler) persistTopUp(ctx context.Context, username string, amount decimal.Decimal, cur models.CurrencyEnum, isCommandCorrect bool, bot *telego.Bot, chatID int64) {
	if !isCommandCorrect {
		return
	}

	err := h.service.TopUpMoney.PersistTopUp(ctx, username, amount, cur)
	if err != nil {
		logrus.Errorf("bank-operation hendler: topUpMoney() : can't top up balance: %s", err.Error())
		if err = h.service.SendMessage(bot, chatID, messages.MsgErrorTopUp); err != nil {
			logrus.Errorf("bank-operation hendler: topUpMoney() : cant't send error TopUp  message %v", err)
			return
		}
	}
	if err = h.service.SendMessage(bot, chatID, messages.MsgTopUpMoney); err != nil {
		logrus.Errorf("bank-operation hendler: topUpMoney() : cant't send TopUp  message %v", err)
	}
}
