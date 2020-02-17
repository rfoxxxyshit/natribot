package main

/* Я если честно в это все срал, извините, куда деваться

Мой первый высер на голанге, поэтому хачю все таки
конструктивной критики, если такая имеется.

Мой TG: @rf0x1d
*/

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	creators := strings.Join(getBotCreators(), ", ")
	log.Printf("Starting natribu.org bot by %s", creators)
	bot, err := tgbotapi.NewBotAPI(getBotToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	_, err = sendMessage(bot, getCreatorID(), "Посылать нахуй готов!")
	if err != nil {
		log.Printf("Напишите че нибудь вашему боту чтобы он мог вам сообщать важную информацию")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start", "go":
			sendMessage(
				bot,
				update.Message.Chat.ID,
				strings.Join(getNatribuOwnMessage(), "\n"),
			)
			sendMessage(
				bot,
				update.Message.Chat.ID,
				strings.Join(getNatribuSecondMessage(), "\n"),
			)
			sendMessage(
				bot,
				update.Message.Chat.ID,
				strings.Join(getNatribuThirdMessage(), "\n"),
			)
		case "source", "repo", "share":
			markup := markupLink{
				"Пойти нахуй!",
				getSourceRepo(),
			}
			replytoMessageWithLink(
				bot,
				update.Message.Chat.ID,
				getSourceMessage(),
				update.Message.MessageID,
				markup,
			)
		case "nahui", "fuck", "goo":
			replytoMessage(
				bot,
				update.Message.Chat.ID,
				"Вы уже тут.",
				update.Message.MessageID,
			)
		}
	}
}
