package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"
	return bot.Send(msg)
}

func replytoMessage(bot *tgbotapi.BotAPI, chatID int64, text string, replyTo int) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyTo
	msg.ParseMode = "HTML"
	return bot.Send(msg)
}

func replytoMessageWithLink(bot *tgbotapi.BotAPI, chatID int64, text string, replyTo int, markup markupLink) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyTo
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = getLinkButton(markup.ButtonText, markup.ButtonLink)
	return bot.Send(msg)
}
