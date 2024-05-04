package main

import (
	"time"

	"github.com/labstack/echo/v4"

	"redoocehub/api/routes"
	"redoocehub/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mysql

	defer app.CloseDBConnection()

	timeout := time.Duration(env.REFRESH_TOKEN_EXPIRY_HOUR) * time.Second

	e := echo.New()

	routes.SetupRoutes(env, timeout, db, e)

	e.Logger.Fatal(e.Start(":8080"))
}
