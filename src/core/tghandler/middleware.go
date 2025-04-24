package tghandler

import (
	"github.com/mymmrac/telego"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"tgtransaction/src/core/tghandler/messages"
)

// Проверка что сумма положительна

func (h *Handler) isCommandCorrect(command string, parts []string, amount decimal.Decimal) bool {

	if ((command == "/send" && len(parts) == 4) || (command == "/topup" && len(parts) == 3)) && amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		logrus.Error("middleware: isCommandCorrect() : wrong amount for command")
		return false
	}
	if command == "/send" && len(parts) != 3 && len(parts) != 4 {
		logrus.Error("middleware: isCommandCorrect() : wrong count arguments in command  /send")
		return false
	}

	if command == "/topup" && len(parts) != 2 && len(parts) != 3 {
		logrus.Error("middleware: isCommandCorrect() : wrong count arguments in command  /top-up")
		return false
	}
	return true
}

func (h *Handler) isRecipientCorrect(recipient string, bot *telego.Bot, chatID int64) bool {
	isExist, err := h.service.UserExistsByUsername(recipient)
	if err != nil {
		logrus.Errorf("middleware: isRecipientCorrect() : error checking user existence: %v", err.Error())
	}
	if !isExist {
		if err = h.service.SendMessage(bot, chatID, messages.MsgUserNotExist); err != nil {
			logrus.Errorf("middleware: isRecipientCorrect() : error sending message UserNotExist: %v", err.Error())
		}
	}
	return isExist

}
