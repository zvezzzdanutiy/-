package main

import (
	"help/internal/api"
	"help/internal/domain/AnekdotProviders"
	"help/internal/domain/TelegramBot"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't read env")
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	anekdotProvider := AnekdotProviders.New(
		&http.Client{},
	)
	apiDomain := api.New(anekdotProvider)
	e.GET(api.GetAnekdotHandlerName, apiDomain.GetAnekdotHandler)

	e.POST(api.AuthHandlerName, apiDomain.AuthHandler)

	e.GET(api.GetTokenName, apiDomain.GetToken)

	e.POST(api.JokeHandlerName, apiDomain.JokeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go e.Start(":" + port)
	tgBot := TelegramBot.New(os.Getenv("BOT_TOKEN"), anekdotProvider)
	tgBot.StartBot()
}
