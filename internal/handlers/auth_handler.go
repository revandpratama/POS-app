package handlers

import (
	"net/http"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/services"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}
type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (h *authHandler) Login(c echo.Context) error {

	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	token, err := h.authService.Login(&req)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    token,
	})

}

func (h *authHandler) Register(c echo.Context) error {
	
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := h.authService.Register(&req); err != nil {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}
