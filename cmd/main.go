package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Start(":3000")

	os.Exit(0)
}
