package ForGeneration

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationDataString(t *testing.T) {
	// Создаем экземпляр AuthorizationData с различными значениями полей
	authData := AuthorizationData{
		FirstName: "John",
		LastName:  "Doe",
		PhotoURL:  "example.com/photo.jpg",
		Username:  "johndoe",
		AuthDate:  1618533961,
		ID:        123,
	}

	// Ожидаемая строка, соответствующая значениям полей authData
	expected := "auth_date=1618533961\nfirst_name=John\nid=123\nlast_name=Doe\nphoto_url=example.com/photo.jpg\nusername=johndoe"

	// Получаем строковое представление authData с помощью метода String()
	result := authData.String()

	// Проверяем, что результат соответствует ожидаемому значению
	assert.Equal(t, expected, result)
}

func TestAuthorizationDataSum(t *testing.T) {
	// Создаем экземпляр AuthorizationData с заданными значениями полей
	authData := AuthorizationData{
		FirstName: "John",
		LastName:  "Doe",
		PhotoURL:  "example.com/photo.jpg",
		Username:  "johndoe",
		AuthDate:  1618533961,
		ID:        123,
	}

	// Токен для теста
	token := "test_token"

	expectedSum := authData.Sum(token)

	// Проверяем, что результат соответствует ожидаемой сумме
	assert.Equal(t, expectedSum, hex.EncodeToString(HashHMAC([]byte(authData.String()), HashSHA256([]byte(token)), sha256.New)))
}
