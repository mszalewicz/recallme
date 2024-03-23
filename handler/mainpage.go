package handler

import (
	body "recallme/view/shared"

	"github.com/labstack/echo/v4"
)

type MainPageHandler struct{}

// func (h *BodyHandler) HandleBodyShow(c echo.Context) error {
// 	return render(c, mainpage.Show())
// }

func (h *MainPageHandler) Show(c echo.Context) error {
	return render(c, body.Page())
}
