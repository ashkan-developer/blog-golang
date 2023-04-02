package Route

import (
	"github.com/labstack/echo/v4"
	"mymodule/Controller"
	"mymodule/Middleware"
)

func Routes(e *echo.Echo) {
	//e.GET("/", func(c echo.Context) error {
	//	return c.JSON(http.StatusOK, "assa")
	//})
	e.POST("/register", Controller.Register)
	e.POST("/login", Controller.Login)

	//	Api Admin	//
	admin := e.Group("/admin", Middleware.Authentication)

	//	Api Category	//
	admin.GET("/all-category", Controller.Category_All)
	admin.POST("/create-category", Controller.Create_category)
	admin.GET("/category/:title", Controller.Category)
	//	Api	User	//
	admin.GET("/users", Controller.Users)
	admin.GET("/user/:user_id", Controller.Userfind)

}
