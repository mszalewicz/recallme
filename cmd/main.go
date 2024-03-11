package main

import (
	"recallme/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	bodyHandler := handler.BodyHandler{}

	app.GET("/", bodyHandler.HandleBodyShow)

	app.Logger.Fatal(app.Start(":3000"))
}
