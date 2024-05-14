package TelegramBot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (d *Domain) SendJoke(userid int64, joke string) error {
	msg := tgbotapi.NewMessage(userid, joke)
	_, err := d.botApi.Send(msg)
	if err != nil {
		return fmt.Errorf("Пользователя не существует")
	} else {
		return nil
	}
}
