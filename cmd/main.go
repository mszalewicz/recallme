package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"recallme/handler"
	server "recallme/model/server"

	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	// Read configuration file

	config := new(server.Config)
	configFile, err := os.ReadFile("../.server_config.toml")

	if err != nil {
		log.Fatal(err)
	}

	err = toml.Unmarshal(configFile, config)

	if err != nil {
		log.Fatal(err)
	}

	// Connect database

	db, err := sql.Open("libsql", config.Database.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", config.Database.Url, err)
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

	os.Exit(1)

	frontHandler := handler.FrontHandler{}
	loginHandler := handler.LoginHandler{}

	app := echo.New()

	app.Static("/static", "static")
	app.GET("/", frontHandler.Show)
	app.GET("/login", loginHandler.Show)

	// Production version
	// app.Logger.Fatal(app.Start(":3000"))

	// Development version
	app.Logger.Fatal(app.Start("127.0.0.1:3000"))
}
