package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"tg_transaction/src/core/repository"
	"tg_transaction/src/core/service"
	"tg_transaction/src/core/tghandler"
	"tg_transaction/src/pkg/pgxhelper"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
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
		log.Fatalf("Не удалось запустить бота: %s", err)
	}

	ctx := context.Background()

	err = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://example.com/bot",
		SecretToken: bot.SecretToken(),
	})

	if err != nil {
		log.Fatalf("Main SetWebhook : не удалось установить настройки webhook: %s", err)
	}

	mux := http.NewServeMux()

	updates, err := bot.UpdatesViaWebhook(ctx, telego.WebhookHTTPServeMux(mux, "/bot", bot.SecretToken()))
	if err != nil {
		log.Fatalf("Main UpdatesViaWebhook: не удалось получить pdates: %s", err)
	}

	go func() {
		_ = http.ListenAndServe(":443", mux)
		log.Print("Сервер запущен")
	}()

	for update := range updates {
		log.Printf("Update: %+v\n", update)
		if update.Message != nil {
			if err = handler.HandleMessage(bot, update.Message); err != nil {
				logrus.Errorf("Main: cant't handle message : %v", err)
			}

		}
	}

}
