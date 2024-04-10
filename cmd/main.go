package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"recallme/handler"
	server "recallme/model/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pelletier/go-toml/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	// Read configuration file

	config := new(server.Config)
	configFile, err := os.ReadFile(".server_config.toml")

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
		log.Fatal(fmt.Sprintf("failed to open db %s: %s", config.Database.Url, err))
	}
	defer db.Close()

	// Logger config - custom format log messages + specified output file

	logFile, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create/open file %s: %s", "logfile", err))
	}

	loggerConfig := middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           logFile,
	}

	// Bootstrap server

	frontHandler := handler.FrontHandler{}
	loginHandler := handler.LoginHandler{DB: db}
	userHandler := handler.UserHandler{DB: db}

	_ = userHandler

	// echo.NewHTTPError(http.Status, message ...interface{})

	app := echo.New()
	app.Use(middleware.LoggerWithConfig(loggerConfig))
	app.Static("/static", "static")
	app.GET("/", frontHandler.Show)
	app.GET("/login", loginHandler.Show)
	app.POST("/signup", loginHandler.SignUp)

	// Production version
	// app.Logger.Fatal(app.Start(":3000"))

	// Development version
	app.Logger.Fatal(app.Start("127.0.0.1:3000"))

}
