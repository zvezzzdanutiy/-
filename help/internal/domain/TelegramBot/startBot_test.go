package TelegramBot

import (
	"testing"
	"time"
)

type MockBotAPI struct{}

func TestStartBot(t *testing.T) {

	d := New("7122476551:AAHZyDfGFku4UjjveplN_hpsxUNFk7W99LA", &MockAnekdotProvider{})

	go d.StartBot()

	time.Sleep(2 * time.Second)

}
