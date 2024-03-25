package handler

import (
	login "recallme/view/login"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct{}

// func (h *BodyHandler) HandleBodyShow(c echo.Context) error {
// 	return render(c, mainpage.Show())
// }

func (h *LoginHandler) Show(c echo.Context) error {
	return render(c, login.Show())
}
