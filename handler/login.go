package handler

import (
	"database/sql"
	"fmt"
	"recallme/model"
	login "recallme/view/login"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	DB *sql.DB
}

func (loginHandler *LoginHandler) Show(c echo.Context) error {
	return render(c, login.Show())
}

func (loginHandler *LoginHandler) SignUp(c echo.Context) error {

	user := model.UserModel{DB: loginHandler.DB}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	emailExists, err := user.EmailExists(email)
	if err != nil {
		return err
	}

	usernameExists, err := user.UsernameExists(username)
	if err != nil {
		return err
	}

	if emailExists {
		return fmt.Errorf("email already exists")
	} else if usernameExists {
		return fmt.Errorf("username already exists")
	} else {
		addedUser, err := user.Add(username, email, password)
		if err != nil {
			return err
		}
		if !addedUser {
			return fmt.Errorf("could not add user")
		}
		return nil
	}

}

//todo login
