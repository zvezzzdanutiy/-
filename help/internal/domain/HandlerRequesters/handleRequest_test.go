package HandlerRequesters

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRequestOne(t *testing.T) {
	e := echo.New()

	// Тест кейсы для HandleRequestTwo
	testCases := []struct {
		requestBody string
		expectedErr string
	}{
		// Test case: Проверяем обработку отсутствующего поля hash
		{`{}`, "FirstName is empty"},
		// Test case: Проверяем обработку некорректного типа данных для auth_date
		{`{"first_name": "John"}`, "Hash is empty"},
		// Test case: Проверяем обработку отсутствующего поля first_name

		{`{"first_name": "John", "hash": "hash_value"}`, "LastName is empty"},

		// Test case: Проверяем обработку некорректного типа данных для last_name

		{`{"first_name": "John", "hash": "hash_value", "last_name": "Doe"}`, "PhotoURL is empty"},

		{`{"first_name": "John", "hash": "hash_value", "last_name": "Doe", "photo_url": "example.com/photo.jpg"}`, "Username is empty"},

		{`{"first_name": "John", "hash": "hash_value", "last_name": "Doe", "photo_url": "example.com/photo.jpg", "username": "johndoe"}`, "AuthDate is empty"},

		{`{"first_name": "John", "hash": "hash_value", "last_name": "Doe", "photo_url": "example.com/photo.jpg", "username": "johndoe", "auth_date": 123}`, "ID is empty"},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		authData, _, err := HandleRequestOne(c)

		if tc.expectedErr != "" {
			// Проверяем, что ошибка не nil и содержит ожидаемое сообщение
			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedErr, err.Error())
			// Проверяем, что authData равен nil
		} else {
			// Проверяем, что ошибки nil
			assert.Nil(t, err)
			// Проверяем, что authData не равен nil
			assert.NotNil(t, authData)
		}
	}

}

func TestHandleRequestTwo(t *testing.T) {
	e := echo.New()

	// Тест кейсы для HandleRequestTwo
	testCases := []struct {
		requestBody string
		expectedErr string
	}{
		// Test case: Проверяем обработку отсутствующего поля chat_id
		{`{}`, "ChatID is empty"},
		// Test case: Проверяем обработку некорректного типа данных для chat_id
		{`{"chat_id": 1231231, "joke": ""}`, "Joke is empty"},
		// Добавьте дополнительные тесты для других полей и их комбинаций значений
		{`{"chat_id": 0, "joke": "Why couldn't the bicycle stand up by itself? Because it was two-tired!"}`, "ChatID is empty"},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tc.requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		_, err := HandleRequestTwo(c)

		assert.Error(t, err)
		assert.Equal(t, tc.expectedErr, err.Error())
	}
}
