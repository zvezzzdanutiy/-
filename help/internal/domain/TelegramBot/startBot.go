package TelegramBot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (d *Domain) StartBot() {
	d.botApi.Debug = true

	log.Printf("Authorized on account %s", d.botApi.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	updates, _ := d.botApi.GetUpdatesChan(u)
	for update := range updates {
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Чтобы получить анекдоты, введите /generate")
			d.botApi.Send(msg)
		} else if update.Message.Text == "/generate" {
			go func(update tgbotapi.Update) {
				anekdot, err := d.anekdotProvider.GetAnekdot(context.Background(), "")
				if err != nil {
					log.Println("Ошибка GetAnekdot")
				}
				msg := tgbotapi.NewMessage(int64(update.Message.From.ID), anekdot)
				d.botApi.Send(msg)
			}(update)
		}
	}
	
}
