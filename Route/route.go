package Route

import (
	"github.com/labstack/echo/v4"
	"mymodule/Controller"
	"mymodule/Middleware"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "assa")
	})
	e.POST("/register", Controller.Register)
	e.POST("/login", Controller.Login)

	//	Api Admin	//
	admin := e.Group("/admin", Middleware.Authentication)

	//	Api	User	//
	admin.GET("/users", Controller.Users)
	admin.GET("/user/:user_id", Controller.Userfind)

	//	Api Category	//
	admin.GET("/all-category", Controller.Category_All)
	admin.POST("/create-category", Controller.Create_category)
	admin.GET("/category/:title", Controller.Category)
	admin.PUT("/category/:title", Controller.UpdateCategory)
	admin.DELETE("/category/:title", Controller.DestroyCategory)

	//	Api Blog	//
	admin.GET("/all-blog", Controller.Blog_All)
	admin.POST("/create-blog", Controller.Create_Blog)

}
