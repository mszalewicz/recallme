package handler

import (
	front "recallme/view/front"

	"github.com/labstack/echo/v4"
)

type FrontHandler struct{}

// func (h *BodyHandler) HandleBodyShow(c echo.Context) error {
// 	return render(c, mainpage.Show())
// }

func (h *FrontHandler) Show(c echo.Context) error {
	return render(c, front.Show())
}
