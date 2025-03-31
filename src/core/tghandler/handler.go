package tghandler

import (
	"github.com/mymmrac/telego"
	"strings"
	"tg_transaction/src/core/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleMessage(bot *telego.Bot, message *telego.Message) error {

	text := message.Text
	chatID := message.Chat.ID

	parts := strings.Split(text, " ")
	switch parts[0] {
	case "/start":
		h.sendStartMessage(bot, chatID)
	case "/help":
		h.sendHelpMessage(bot, chatID)
	case "/send":
		h.sendMoney(bot, chatID, parts[1], parts[2])
	case "/top-up":
		h.topUpMoney(bot, chatID, parts[1])
	default:
		h.sendUnknownMessage(bot, chatID)
	}

	return nil
}
