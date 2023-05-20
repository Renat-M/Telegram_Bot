package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Посмеяться"),
		tgbotapi.NewKeyboardButton("Кнопка 2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кнопка 3"),
		tgbotapi.NewKeyboardButton("Кнопка 4"),
	),
)

var laughingKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Анекдот"),
		tgbotapi.NewKeyboardButton("Мемчик"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var keyboard2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кнопка 2.1"),
		tgbotapi.NewKeyboardButton("Кнопка 2.2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var keyboard3 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кнопка 3.1"),
		tgbotapi.NewKeyboardButton("Кнопка 3.2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var keyboard4 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Кнопка 4.1"),
		tgbotapi.NewKeyboardButton("Кнопка 4.2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)
