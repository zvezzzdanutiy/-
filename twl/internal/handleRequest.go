package internal

import (
	"log"

	"github.com/labstack/echo/v4"
)

func HandleRequest(c echo.Context) (*AuthorizationData, error) {
	req := new(AuthorizationData)
	if err := c.Bind(req); err != nil {
		log.Println("Ошибка присваивания")
		return nil, err
	}
	return req, nil
}
