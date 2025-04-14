package tghandler

import (
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"tg_transaction/src/core/tghandler/messages"
)

func (h *Handler) persistTransaction(username string, amount float64, recipient string, isCommandCorrect bool, bot *telego.Bot, chatID int64) {

	if !h.isRecipientCorrect(recipient, bot, chatID) {
		return
	}
	if !isCommandCorrect {
		return
	}
	err := h.service.Transaction.PersistTransaction(username, amount, recipient)
	if err != nil {
		logrus.Errorf("bank-operation hendler: persistTransaction() :cant't send money: %s", err.Error())
		if sendErr := h.service.SendMessage(bot, chatID, messages.MsgErrorSend); sendErr != nil {
			logrus.Errorf("bank-operation hendler: persistTransaction(): cant't send error Send  message %v", sendErr)

		}
		return
	}
	if err := h.service.SendMessage(bot, chatID, messages.MsgSendMoney); err != nil {
		logrus.Errorf("bank-operation hendler: persistTransaction() :cant't send Send message %v", err)
	}
}

func (h *Handler) persistTopUp(username string, amount float64, isCommandCorrect bool, bot *telego.Bot, chatID int64) {
	if !isCommandCorrect {
		return
	}
	err := h.service.TopUpMoney.PersistTopUp(username, amount)
	if err != nil {
		logrus.Errorf("bank-operation hendler: topUpMoney() : can't top up balance: %s", err.Error())
		if err = h.service.SendMessage(bot, chatID, messages.MsgErrorTopUp); err != nil {
			logrus.Errorf("bank-operation hendler: topUpMoney() : cant't send error TopUp  message %v", err)
		}
	}
	if err = h.service.SendMessage(bot, chatID, messages.MsgTopUpMoney); err != nil {
		logrus.Errorf("bank-operation hendler: topUpMoney() : cant't send TopUp  message %v", err)
	}
}

func (h *Handler) takeBalance(username string) float64 {
	balance, err := h.service.TakeTotalBalance(username)
	if err != nil {
		logrus.Errorf("bank-operation hendler: takeBalance() : Error take balance: %s", err.Error())

	}
	return balance
}
