package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	InitialKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Пройти тест"),
		),
	)
	GreetingKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Старт"),
			tgbotapi.NewKeyboardButton("Выход"),
		),
	)
	BoolKeyBoard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Да"),
			tgbotapi.NewKeyboardButton("Нет"),
			tgbotapi.NewKeyboardButton("Выход"),
		),
	)
	TestListKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Тест \"Честность\""),
		),
	)
)
