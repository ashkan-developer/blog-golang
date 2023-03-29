package Middleware

import (
	"github.com/labstack/echo/v4"
	"mymodule/Helper"
	"net/http"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientToken := c.Request().Header.Get("token")

		if clientToken == "" {
			res := c.JSON(http.StatusInternalServerError, "No Authorization header provided")
			return res
		}

		claims, err := Helper.ValidateToken(clientToken)

		if err != "" {
			res := c.JSON(http.StatusInternalServerError, "error")
			return res
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		return next(c)
	}

}
