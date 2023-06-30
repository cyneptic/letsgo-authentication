package main

import (
	"github.com/cyneptic/letsgo-authentication/controller"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	_ = godotenv.Load(".env")
	e := echo.New()
	controller.AddAuthServiceRoutes(*e)
	log.Fatal(e.Start(":8080"))
}
