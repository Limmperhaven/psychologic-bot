package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"psycho-tg-bot/internal/telegram"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error occured loading env: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("error occured creating bot api: %s", err.Error())
	}

	bot.Debug = false

	telegramBot := telegram.NewBot(bot)

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
