package main

import (
	"fmt"
	"log"
	"os"
	"url-shortner/api/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/:url", routes.ResolveURL)
	e.POST("/api/v1", routes.ShortenURL)

	log.Fatal(e.Start(os.Getenv("APP_PORT")))
}
