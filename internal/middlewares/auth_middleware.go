package middlewares

import (
	"net/http"
	"point-of-sales-app/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization format")
		}

		token := parts[1]
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		claims, err := helper.VerifyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		c.Set("user_id", claims.ID)
		c.Set("user_name", claims.Name)
		c.Set("user_email", claims.Email)

		return next(c)
	}
}
