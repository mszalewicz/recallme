package handler

import (
	"recallme/view/mainpage"

	"github.com/labstack/echo/v4"
)

type BodyHandler struct{}

func (h BodyHandler) HandleBodyShow(c echo.Context) error {
	return render(c, mainpage.Show())
}
