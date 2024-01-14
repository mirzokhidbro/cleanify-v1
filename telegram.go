package main

import (
	"bw-erp/internal/app"
	"bw-erp/internal/bot_service"
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	app.RunMigration()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(bot_service.Handler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if nil != err {
		panic(err)
	}

	go func() {
		http.ListenAndServe(":8080", b.WebhookHandler())
	}()

	b.StartWebhook(ctx)

}
