package routes

import (
	"point-of-sales-app/adapter"
	"point-of-sales-app/internal/handlers"
	"point-of-sales-app/internal/middlewares"
	"point-of-sales-app/internal/repositories"
	"point-of-sales-app/internal/services"

	"github.com/labstack/echo/v4"
)

func initProductHandler() handlers.ProductHandler {
	repo := repositories.NewProductRepository(adapter.DB)
	service := services.NewProductService(repo)
	return handlers.NewProductHandler(service)
}

func InitProductRoutes(e *echo.Group) {
	productHandler := initProductHandler()

	e.Use(middlewares.AuthMiddleware)
	
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProduct)
	e.POST("/products", productHandler.CreateProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)
}
