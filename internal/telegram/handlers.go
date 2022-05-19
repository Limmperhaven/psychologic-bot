package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"psycho-tg-bot/internal/keyboards"
	"psycho-tg-bot/internal/tests"
	"strings"
)

const (
	commandStart    = "start"
	commandHelp     = "help"
	commandTestList = "testlist"
)

func (b *Bot) handleCommand(msg *tgbotapi.Message) error {
	message := tgbotapi.NewMessage(msg.Chat.ID, "")
	switch msg.Command() {
	case commandStart:
		message.Text = fmt.Sprintf("Привет, %s %s! Я психологический бот, могу показать тебе несколько психологических "+
			"тестов! Если интересно, жми \"Пройти тест\"!", msg.From.FirstName, msg.From.LastName)
		message.ReplyMarkup = keyboards.InitialKeyboard
	case commandHelp:
		message.Text = fmt.Sprintf("Привет, %s %s! Я психологический бот, могу показать тебе несколько психологических "+
			"тестов! Если интересно, жми \"Пройти тест\"!", msg.From.FirstName, msg.From.LastName)
		message.ReplyMarkup = keyboards.InitialKeyboard
	case commandTestList:
		message.Text = fmt.Sprintf("Сейчас я знаю следующие тесты: \n\n\t 1.Тест на честность")
		message.ReplyMarkup = keyboards.TestListKeyboard
	default:
		message.Text = "Я не знаю такой команды(("
	}
	_, err := b.Bot.Send(message)
	return err
}

func (b *Bot) handleMessage(msg *tgbotapi.Message) error {
	if b.TestChannel[msg.Chat.ID] != nil {
		select {
		case b.TestChannel[msg.Chat.ID] <- *msg:
			testResp := <-b.TestChannel[msg.Chat.ID]
			hResp, ok := testResp.(tgbotapi.MessageConfig)
			if !ok {
				return errors.New("invalid response from test")
			}
			_, err := b.Bot.Send(hResp)
			return err
		case <-b.TestChannel[msg.Chat.ID]:
			b.TestChannel[msg.Chat.ID] = nil
		}
	}

	switch strings.ToLower(msg.Text) {
	case "тест \"честность\"":
		b.TestChannel[msg.Chat.ID] = tests.StartHonestyTest()
		_, err := b.Bot.Send(tests.GetHonestyGreeting(msg.Chat.ID))
		return err
	case "пройти тест":
		message := tgbotapi.NewMessage(msg.Chat.ID, "")
		message.Text = fmt.Sprintf("Сейчас я знаю следующие тесты: \n\n\t 1.Тест на честность")
		message.ReplyMarkup = keyboards.TestListKeyboard
		_, err := b.Bot.Send(message)
		return err
	default:
		_, err := b.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Я не понимаю, что ты хочешь(("))
		return err
	}
}
