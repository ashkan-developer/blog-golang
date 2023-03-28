package main

import (
	"github.com/labstack/echo/v4"
	"mymodule/Route"
)

func main() {

	server := echo.New()
	Route.Routes(server)

	server.Logger.Fatal(server.Start(":8000"))

}
