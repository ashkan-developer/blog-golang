package Route

import (
	"github.com/labstack/echo/v4"
	"mymodule/Controller"
)

func Routes(e *echo.Echo) {
	e.POST("/register", Controller.Register)
}
