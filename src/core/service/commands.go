package service

import (
	"context"
	"fmt"
	"github.com/mymmrac/telego"
	"tg_transaction/src/core/repository"
)

type CommandService struct {
	repo repository.Commands
}

func NewCommandService(repo repository.Commands) *CommandService {
	return &CommandService{}
}

func (c *CommandService) SendMoney(bot *telego.Bot, chatID int64, amount, recipient string) error {
	return c.repo.SendMoney(bot, chatID, amount, recipient)
}

func (c *CommandService) TopUpMoney(bot *telego.Bot, chatID int64, amount string) error {
	return c.repo.TopUpMoney(bot, chatID, amount)
}

func (c *CommandService) SendMessage(bot *telego.Bot, chatID int64, msg string) error {
	Id := telego.ChatID{ID: chatID}
	_, err := bot.SendMessage(context.Background(), &telego.SendMessageParams{
		ChatID: Id,
		Text:   msg,
	})
	if err != nil {
		return fmt.Errorf("handler:failed to send message: %w", err)
	}
	return nil
}
