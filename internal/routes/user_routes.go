package routes

import (
	"point-of-sales-app/adapter"
	"point-of-sales-app/internal/handlers"
	"point-of-sales-app/internal/middlewares"
	"point-of-sales-app/internal/repositories"
	"point-of-sales-app/internal/services"

	"github.com/labstack/echo/v4"
)

func initUserHandler() handlers.UserHandler {
	repo := repositories.NewUserRepository(adapter.DB)
	service := services.NewUserService(repo)
	return handlers.NewUserHandler(service)
}

func InitUserRoutes(e *echo.Group) {
	handler := initUserHandler()

	e.Use(middlewares.AuthMiddleware)
	
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}
