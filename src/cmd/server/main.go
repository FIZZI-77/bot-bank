package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"tg_transaction/src/core/repository"
	"tg_transaction/src/core/service"
	"tg_transaction/src/core/tghandler"
	"tg_transaction/src/pkg/pgxhelper"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}

	db, err := pgxhelper.NewPostgresDB(pgxhelper.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})

	if err != nil {
		log.Fatalf("Ошибка подключения к базе данны: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := tghandler.NewHandler(services)

	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		log.Fatalf("Не удалось запустить бота: %s", err.Error())
	}

	ctx := context.Background()

	offset := 0

	var updates []telego.Update
	for {

		updates, err = bot.GetUpdates(ctx, &telego.GetUpdatesParams{
			Offset:  offset,
			Timeout: 8,
		})
		if err != nil {
			log.Fatalf("Ошибка получения обновлений: %s", err.Error())
		}

		for _, update := range updates {

			log.Printf("Пришло обновление: %+v\n", update)
			if update.Message.Text != "" {
				err = handler.HandleMessage(bot, update.Message)
				if err != nil {
					logrus.Errorf("Ошибка обработки сообщения: %v", err)
				}
			}

			offset = update.UpdateID + 1
		}

	}

}
