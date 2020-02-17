package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func getLinkButton(text string, link string) tgbotapi.InlineKeyboardMarkup {
	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(text, link),
		),
	)
	return markup
}
