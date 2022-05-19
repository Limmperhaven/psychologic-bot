package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	Bot         *tgbotapi.BotAPI
	TestChannel map[int64]chan interface{}
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		Bot:         bot,
		TestChannel: make(map[int64]chan interface{}),
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.Bot.Self.UserName)

	updates := b.initUpdatesChannel()

	return b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommand(update.Message)
			if err != nil {
				return err
			}
			continue
		}

		err := b.handleMessage(update.Message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updatesChannel, err := b.Bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		log.Fatalf("error generation updates channel: %s", err.Error())
	}

	return updatesChannel
}
