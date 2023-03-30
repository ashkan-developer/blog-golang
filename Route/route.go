package Route

import (
	"github.com/labstack/echo/v4"
	"mymodule/Controller"
	"mymodule/Middleware"
)

func Routes(e *echo.Echo) {
	e.POST("/register", Controller.Register)
	e.POST("/login", Controller.Login)

	admin := e.Group("/admin", Middleware.Authentication)

	admin.GET("/all-category", Controller.Category_All)
	admin.POST("/create-category", Controller.Create_category)
	admin.GET("/users", Controller.Users)

}
