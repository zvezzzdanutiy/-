package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"twl/internal"

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

	// Обработчик POST-запроса для /answer

	e.POST("/auth", func(c echo.Context) error {
		req, err := internal.HandleRequest(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}
		fmt.Println(req.Sum(os.Getenv("BOT_TOKEN")) == req.Hash)

		if req.Sum(os.Getenv("BOT_TOKEN")) == req.Hash {
			return c.JSON(http.StatusOK, true)
		} else {
			return c.JSON(http.StatusOK, false)
		}
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)

}
