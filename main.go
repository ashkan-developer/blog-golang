package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mymodule/Route"
)

func main() {

	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	Route.Routes(server)

	server.Logger.Fatal(server.Start(":8000"))

}
