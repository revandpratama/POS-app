package routes

import (
	"point-of-sales-app/adapter"
	"point-of-sales-app/internal/handlers"
	"point-of-sales-app/internal/repositories"
	"point-of-sales-app/internal/services"

	"github.com/labstack/echo/v4"
)


func initHandler() handlers.AuthHandler {
	repo := repositories.NewAuthRepository(adapter.DB)
	service := services.NewAuthService(repo)
	return handlers.NewAuthHandler(service)
}

func InitAuthRoutes(e *echo.Group) {
	handler := initHandler()
	
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}
