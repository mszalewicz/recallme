package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"recallme/handler"

	"github.com/labstack/echo/v4"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	// TODO: Read server configuration file

	// TODO: Assign url from the consiguration file

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	query, err := db.Query("SELECT * FROM users LIMIT 2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	type Users struct {
		id   int
		name string
	}

	user := new(Users)

	for query.Next() {
		err := query.Scan(&user.id, &user.name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
	}

	mainpageHandler := handler.MainPageHandler{}
	// loginHandler := handler.LoginPageHandler{}

	app := echo.New()

	app.Static("/static", "static")
	app.GET("/", mainpageHandler.Show)
	// app.GET("/login", loginHandler.Show)

	app.Logger.Fatal(app.Start(":3000"))
}
