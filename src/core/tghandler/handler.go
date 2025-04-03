package tghandler

import (
	"github.com/mymmrac/telego"
	"log"
	"strconv"
	"strings"
	"tg_transaction/src/core/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleMessage(bot *telego.Bot, message *telego.Message) (err error) {

	text := message.Text
	chatID := message.Chat.ID
	username := message.Chat.Username
	telegramID := message.From.ID
	parts := strings.Split(text, " ")

	var amount float64
	if len(parts) != 1 {
		amount, err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatalf("Error converting amount to int: %v", err)
		}
		h.checkCommand(parts[0], parts, amount)
	}

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
	case parts[0] == "/topup" && len(parts) == 2:
		h.topUpMoney(username, amount)

	case parts[0] == "/send" && len(parts) == 3:
		h.sendMoney(username, amount, parts[2])
	case parts[0] == "/balance" && len(parts) == 1:
		h.takeBalanceMessage(bot, chatID, username)
	default:
		h.sendUnknownMessage(bot, chatID)
	}

	return nil
}
