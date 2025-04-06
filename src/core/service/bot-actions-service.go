package service

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
)

type BotActionsService struct {
}

func NewBotActionsService() *BotActionsService {
	return &BotActionsService{}
}

func (c *BotActionsService) SendMessage(bot *telego.Bot, chatID int64, msg string) error {
	Id := telego.ChatID{ID: chatID}
	_, err := bot.SendMessage(context.Background(), &telego.SendMessageParams{
		ChatID: Id,
		Text:   msg,
	})
	if err != nil {
		return fmt.Errorf("bot-action service: SendMessage() :failed to send message: %w", err)
	}
	return nil
}
