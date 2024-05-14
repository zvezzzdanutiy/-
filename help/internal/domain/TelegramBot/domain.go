package TelegramBot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Domain struct {
	botApi          *tgbotapi.BotAPI
	anekdotProvider AnekdotProvider
}

func New(token string, anekdotProvider AnekdotProvider) *Domain {
	s, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Ошибка tgbotapi")
	}
	return &Domain{
		botApi:          s,
		anekdotProvider: anekdotProvider,
	}
}
