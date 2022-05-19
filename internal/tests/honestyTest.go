package tests

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"psycho-tg-bot/internal/keyboards"
)

func StartHonestyTest() chan interface{} {

	mc := make(chan interface{})

	go func() {

		data, err := ioutil.ReadFile("internal/tests/questions.json")
		if err != nil {
			log.Fatalf(err.Error())
		}

		questionList := gjson.Get(string(data), "honesty").Array()
		answerList := make([]uint8, 0)
		folQuestion := -1

		for {
			request := <-mc

			msg, ok := request.(tgbotapi.Message)
			if !ok {
				log.Fatalf("Invalid message type")
			}

			if folQuestion == -1 {
				if msg.Text == "Выход" {
					response := tgbotapi.NewMessage(msg.Chat.ID, "Выход из теста...")
					response.ReplyMarkup = keyboards.InitialKeyboard
					mc <- response
					break
				} else {
					folQuestion++
				}
			} else {
				switch msg.Text {
				case "Да":
					answerList = append(answerList, 1)
					folQuestion++
				case "Нет":
					answerList = append(answerList, 0)
					folQuestion++
				case "Выход":
					response := tgbotapi.NewMessage(msg.Chat.ID, "Выход из теста...")
					response.ReplyMarkup = keyboards.InitialKeyboard
					mc <- response
					break
				default:
					mc <- tgbotapi.NewMessage(msg.Chat.ID, "Некорректный ответ")
					continue
				}
			}

			response := tgbotapi.NewMessage(msg.Chat.ID, "")

			if folQuestion != len(questionList) {
				response.Text = questionList[folQuestion].String()
				response.ReplyMarkup = keyboards.BoolKeyBoard
				mc <- response
			} else {
				response = CalculateScore(answerList, msg.Chat.ID)
				mc <- response
				break
			}
		}

		mc <- "quit"

	}()

	return mc
}

func GetHonestyGreeting(chatID int64) tgbotapi.MessageConfig {
	data, err := ioutil.ReadFile("internal/tests/greetings.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	greeting := tgbotapi.NewMessage(chatID, gjson.Get(string(data), "honesty").String())
	greeting.ReplyMarkup = keyboards.GreetingKeyboard

	return greeting
}

func CalculateScore(answerList []uint8, chatID int64) tgbotapi.MessageConfig {
	score := answerList[0] + answerList[2] + answerList[4] + answerList[5] + answerList[7] + answerList[8] + answerList[9] +
		+answerList[10] + answerList[11] + answerList[13] + answerList[14] + answerList[15] + answerList[16] +
		+answerList[17] + answerList[18] + answerList[19] + answerList[20] + answerList[21] + answerList[22] +
		+answerList[23] + answerList[24] + answerList[27] + answerList[28] + answerList[29] + answerList[30] +
		+answerList[31] + answerList[32] + answerList[33] + 6 - answerList[1] - answerList[3] - answerList[6] - answerList[12] -
		-answerList[25] - answerList[26]
	response := tgbotapi.NewMessage(chatID, "Ошибка при вычислении результата")
	response.ReplyMarkup = keyboards.InitialKeyboard

	switch {
	case score <= 5:
		response.Text = "Результат:\n\nОчень низкий показатель по шкале \"Честность\". Свидетельствует о ярко выраженной склонности ко лжи, приукрашиванию себя. Также может свидетельствовать о низких показателях социального интеллекта."
		return response
	case score > 5 && score <= 13:
		response.Text = "Результат:\n\nНизкий показатель по шкале \"Честность\". Свидетельствует о значительной склонности ко лжи. Любит приукрашивать себя, своё поведение."
		return response
	case score > 13 && score <= 29:
		response.Text = "Результат:\n\nНормальный результат. Склонность ко лжи не выявлена. Может быть, изредка склонен приукрашивать себя, своё поведение, но в пределах нормы."
		return response
	case score > 29:
		response.Text = "Результат:\n\nВысокий результат по шкале \"Честность\". Такой высокий результат может быть связан не только с высокой личностной честностью, но и следствием других причин: преднамеренного искажения ответов, очень неверной самооценки. Следует осторожно отнестись к данному результату."
		return response
	default:
		return response
	}
}
