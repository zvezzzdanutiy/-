package HandlerRequesters

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"help/internal/domain/ForGeneration"
)

func HandleRequestOne(c echo.Context) (*ForGeneration.AuthorizationData, error, error) {
	// Создаем новый экземпляр AuthorizationData
	req := new(ForGeneration.AuthorizationData)
	// Привязываем данные из запроса к структуре AuthorizationData
	if err := c.Bind(req); err != nil {
		return nil, err, nil
	}
	return req, nil, checkEmptyFieldsRequestersOne(*req)
}

func HandleRequestTwo(c echo.Context) (*ForGeneration.Data, error) {
	req := new(ForGeneration.Data)
	// Привязываем данные из запроса к структуре AuthorizationData
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	return req, checkEmptyFieldsRequestersTwo(*req)
}

func checkEmptyFieldsRequestersOne(data ForGeneration.AuthorizationData) error {
	// Проверяем каждое поле отдельно
	if data.FirstName == "" {
		return fmt.Errorf("FirstName is empty")
	} else if data.Hash == "" {
		return fmt.Errorf("Hash is empty")
	} else if data.LastName == "" {
		return fmt.Errorf("LastName is empty")
	} else if data.PhotoURL == "" {
		return fmt.Errorf("PhotoURL is empty")
	} else if data.Username == "" {
		return fmt.Errorf("Username is empty")
	} else if data.AuthDate == 0 {
		return fmt.Errorf("AuthDate is empty")
	} else if data.ID == 0 {
		return fmt.Errorf("ID is empty")
	} else {
		return nil
	}
}
func checkEmptyFieldsRequestersTwo(data ForGeneration.Data) error {

	// Проверяем каждое поле отдельно
	if data.ChatId == 0 {
		return fmt.Errorf("ChatID is empty")
	} else if data.Joke == "" {
		return fmt.Errorf("Joke is empty")
	} else {
		return nil
	}
}
