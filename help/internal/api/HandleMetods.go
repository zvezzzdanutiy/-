package api

import (
	"fmt"
	"help/internal/domain/AnekdotProviders"
	"help/internal/domain/HandlerRequesters"
	"help/internal/dto"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

const GetAnekdotHandlerName = "/anekdot"
const AuthHandlerName = "/auth"
const JokeHandlerName = "/send-joke"
const GetTokenName = "/token"

func (d *Domain) GetAnekdotHandler(c echo.Context) error {
	AnekdotProviders := &AnekdotProviders.Domain{}
	anekdot, err := d.anekdotProvider.GetAnekdot(c.Request().Context(), AnekdotProviders.GenerateURL())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			ErrorMessage: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Anekdot{
		Text: anekdot,
	})
}

func (d *Domain) AuthHandler(c echo.Context) error {
	req, err, _ := HandlerRequesters.HandleRequestOne(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	if req.Sum(os.Getenv("BOT_TOKEN")) == req.Hash {
		return c.JSON(http.StatusOK, true)
	} else {
		return c.JSON(http.StatusOK, false)
	}
}

func (d *Domain) JokeHandler(c echo.Context) error {
	req, err := HandlerRequesters.HandleRequestTwo(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	// Ваш код для отправки анекдота в Telegram
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", os.Getenv("BOT_TOKEN"))
	messageData := url.Values{}
	messageData.Set("chat_id", strconv.FormatInt(req.ChatId, 10))
	messageData.Set("text", req.Joke)

	resp, err := http.PostForm(telegramURL, messageData)
	defer resp.Body.Close()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to send joke")
	}

	return c.JSON(http.StatusOK, "Joke sent successfully")
}

func (d *Domain) GetToken(c echo.Context) error {
	return c.String(http.StatusOK, os.Getenv("BOT_TOKEN"))
}
