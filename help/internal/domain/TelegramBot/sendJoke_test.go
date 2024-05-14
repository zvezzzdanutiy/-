package TelegramBot

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAnekdotProvider struct{}

func (m *MockAnekdotProvider) GetAnekdot(ctx context.Context, category string) (string, error) {
	// Возвращаем заранее определенный анекдот для теста
	return "Тестовый анекдот", nil
}

func TestSendJoke(t *testing.T) {
	// Создаем экземпляр Domain с MockAnekdotProvider
	domain := New("7122476551:AAHZyDfGFku4UjjveplN_hpsxUNFk7W99LA", &MockAnekdotProvider{})

	// Проверяем, что sendJoke возвращает ошибку, если передается пустой ID пользователя
	err := domain.SendJoke(12312312312, "Тестовый анекдот")
	assert.NotNil(t, err)

	// Проверяем, что sendJoke возвращает nil, если все условия выполнены успешно
	err = domain.SendJoke(920849123, "Тестовый анекдот")
	assert.Nil(t, err)
}
