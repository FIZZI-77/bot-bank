package tghandler

import (
	"context"
	"github.com/mymmrac/telego"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
	"tgtransaction/src/core/models"
	"tgtransaction/src/core/service"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleMessage(bot *telego.Bot, message *telego.Message) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	text := message.Text
	chatID := message.Chat.ID
	username := message.Chat.Username
	telegramID := message.From.ID
	parts := strings.Split(text, " ")

	var amount decimal.Decimal
	if len(parts) != 1 {
		amount, err = decimal.NewFromString(parts[1])
		if err != nil {
			log.Fatalf("Handler : HandleMessage(): Error converting amount to int: %v", err)
		}
	}

	var cur models.CurrencyEnum
	if len(parts) > 2 {
		cur, err = ParseCurrencyToEnum(strings.ToUpper(parts[2]))

		if err != nil {
			logrus.Error("Handler : HandleMessage(): Error converting currency to enum: %v", err)
			h.sendErrorCurrencyMessage(bot, chatID)
			return
		}
	}
	isCommandCorrect := h.isCommandCorrect(parts[0], parts, amount)

	switch {
	case parts[0] == "/start" && len(parts) == 1:
		if err := h.sendStartMessage(bot, chatID); err != nil {
			return err
		}
		h.addUser(username, telegramID)
	case parts[0] == "/help" && len(parts) == 1:
		if err := h.sendHelpMessage(bot, chatID); err != nil {
			return err
		}
	case parts[0] == "/topup" && len(parts) == 3:
		h.persistTopUp(ctx, username, amount, cur, isCommandCorrect, bot, chatID)
	case parts[0] == "/send" && len(parts) == 4:
		h.persistTransaction(ctx, username, amount, parts[3], isCommandCorrect, bot, chatID, cur)
	case parts[0] == "/balance" && len(parts) == 1:
		h.takeBalanceMessage(ctx, bot, chatID, username)
	case parts[0] == "/topup" && len(parts) == 2:
		h.sendErrorCurrencyMessage(bot, chatID)
	default:
		h.sendUnknownMessage(bot, chatID)
	}

	return nil
}
