package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"tgtransaction/src/core/repository"
	"tgtransaction/src/core/service"
	"tgtransaction/src/core/tghandler"
	"tgtransaction/src/pkg/pgxhelper"
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
		log.Fatalf("Error connect to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := tghandler.NewHandler(services)

	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		log.Fatalf("can't create new bot: %s", err.Error())
	}

	params := &telego.DeleteWebhookParams{}
	err = bot.DeleteWebhook(context.Background(), params)

	if err != nil {
		log.Fatalf("can't start bot: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		<-sigs
		cancel()
	}()

	go func() {
		defer wg.Done()
		tgRun(ctx, bot, handler)
	}()

	wg.Wait()

}

func tgRun(ctx context.Context, bot *telego.Bot, handler *tghandler.Handler) {

	offset := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down tgRun...")
			return
		default:
			updates, err := bot.GetUpdates(ctx, &telego.GetUpdatesParams{
				Offset:  offset,
				Timeout: 8,
			})
			if err != nil {
				log.Fatalf("error take updates: %s", err.Error())
			}

			for _, update := range updates {

				log.Printf("Update come: %+v\n", update)
				if update.Message.Text != "" {
					err = handler.HandleMessage(bot, update.Message)
					if err != nil {
						logrus.Errorf("Message processing error: %v", err)
					}
				}

				offset = update.UpdateID + 1
			}
		}
	}
}
