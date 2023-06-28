package main

import (
	"log"
	"github.com/cyneptic/letsgo-authentication/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.AddAuthServiceRoutes(*e)
	log.Fatal(e.Start(":8080"))
}
