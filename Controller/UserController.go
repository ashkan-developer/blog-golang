package Controller

import (
	"github.com/labstack/echo/v4"
	"mymodule/Model"
	"net/http"
)

func Register(c echo.Context) error {

	//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//defer cancel()
	return c.JSON(http.StatusOK, user)

}
