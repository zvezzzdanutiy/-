package TelegramBot

import "testing"

func TestNew(t *testing.T) {
	// Создаем экземпляр домена
	domain := New("7122476551:AAHZyDfGFku4UjjveplN_hpsxUNFk7W99LA", &MockAnekdotProvider{})

	// Проверяем, что экземпляр домена создается без ошибок
	if domain == nil {
		t.Error("Ошибка: экземпляр домена не был создан")
	}

	// Проверяем, что бот API установлен
	if domain.botApi == nil {
		t.Error("Ошибка: бот API не был установлен")
	}

	// Проверяем, что провайдер анекдотов установлен
	if domain.anekdotProvider == nil {
		t.Error("Ошибка: провайдер анекдотов не был установлен")
	}
}
