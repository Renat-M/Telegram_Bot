package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const commandStart = "start"

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды :(")
	switch message.Command() {
	case commandStart:
		msg.Text = "Привет, " + message.From.FirstName + "!"
		msg.ReplyMarkup = startKeyboard
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	msg := tgbotapi.NewMessage(message.Chat.ID, " ")
	switch message.Text {
	case "Посмеяться":
		msg.Text = "Выберите предпочтительный способ посмеяться"
		msg.ReplyMarkup = laughingKeyboard
	case "Кнопка 2":
		msg.Text = "Кнопка 2 - прекрасный выбор!"
		msg.ReplyMarkup = keyboard2
	case "Кнопка 3":
		msg.Text = "Кнопка 3 - прекрасный выбор!"
		msg.ReplyMarkup = keyboard3
	case "Кнопка 4":
		msg.Text = "Кнопка 4 - прекрасный выбор!"
		msg.ReplyMarkup = keyboard4
	case "Назад":
		msg.Text = "Нажми на кнопку, получишь результат"
		msg.ReplyMarkup = startKeyboard
	case "Анекдот":
		msg.Text = GetAnekdot()
	case "Мемчик":
		msg.Text = GetMeme()
	}

	b.bot.Send(msg)
}
