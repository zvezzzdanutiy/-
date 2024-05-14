package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"help/internal/dto"
)

type mockAnekdotProvider struct{}

func (m *mockAnekdotProvider) GetAnekdot(ctx context.Context, category string) (string, error) {
	// Возвращаем фиктивный анекдот и ошибку для тестирования
	return "Mocked anekdot", nil
}

func TestGetAnekdotHandler(t *testing.T) {
	// Инициализируем Echo framework
	e := echo.New()

	// Создаем экземпляр mockAnekdotProvider
	anekdotProvider := &mockAnekdotProvider{}

	// Создаем экземпляр Domain с mockAnekdotProvider
	domain := New(anekdotProvider)

	// Создаем HTTP запрос (GET /anekdot)
	req := httptest.NewRequest(http.MethodGet, "/anekdot", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Вызываем метод GetAnekdotHandler
	if assert.NoError(t, domain.GetAnekdotHandler(c)) {
		// Проверяем, что статус код ответа 200 OK
		assert.Equal(t, http.StatusOK, rec.Code)

		// Проверяем, что ответ содержит ожидаемый анекдот
		expectedResponse := dto.Anekdot{Text: "Mocked anekdot"}
		assert.JSONEq(t, `{"Text":"Mocked anekdot"}`, rec.Body.String())

		// Преобразуем JSON ответ в объект dto.Anekdot для сравнения
		var actualResponse dto.Anekdot
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &actualResponse)) {
			assert.Equal(t, expectedResponse, actualResponse)
		}
	}
}
func TestAuthHandler(t *testing.T) {
	// Создание экземпляра Echo
	e := echo.New()
	d := &Domain{}
	// Регистрация обработчика маршрута
	e.POST(AuthHandlerName, d.AuthHandler)

	// Создание фейкового контекста HTTP запроса
	reqBody := bytes.NewBufferString(`{"key": "value"}`) // Замените на ваше тестовое тело запроса
	req := httptest.NewRequest(http.MethodPost, AuthHandlerName, reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Вызов обработчика через Echo
	if assert.NoError(t, d.AuthHandler(c)) {
		// Проверка кода HTTP ответа
		assert.Equal(t, http.StatusOK, rec.Code)

		// Проверка содержимого тела HTTP ответа
		expectedResponse := "false"
		assert.Equal(t, expectedResponse, strings.TrimSpace(rec.Body.String()))
	}
}

func TestGetToken(t *testing.T) {
	e := echo.New()
	d := &Domain{}
	e.GET(GetTokenName, d.GetToken)
	req := httptest.NewRequest(http.MethodGet, GetTokenName, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, d.GetToken(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, os.Getenv("BOT_TOKEN"), rec.Body.String())
	}
}

// создай unit-тест для моей функции New
func TestNew(t *testing.T) {
	// Создаем мок AnekdotProvider
	anekdotProvider := &mockAnekdotProvider{}

	// Создаем новый экземпляр Domain с моком AnekdotProvider
	domain := New(anekdotProvider)

	assert.Equal(t, domain.anekdotProvider, anekdotProvider)
	if domain.anekdotProvider != anekdotProvider {
		t.Errorf("Ожидалось, что anekdotProvider будет установлен правильно, получено: %v", domain.anekdotProvider)
	}
}

func TestJokeHandler(t *testing.T) {
	e := echo.New()
	d := &Domain{}
	e.POST(JokeHandlerName, d.JokeHandler)
	req := httptest.NewRequest(http.MethodPost, JokeHandlerName, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, d.JokeHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
