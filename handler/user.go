package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	DB *sql.DB
}

func (u UserHandler) SignUpForm(c echo.Context) error {

	// user := new(model.UserModel)
	// err := user.SingUp(db)

	// return err
	return nil
}
