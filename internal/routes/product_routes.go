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

	product := e.Group("/products")
	product.Use(middlewares.AuthMiddleware)

	product.GET("", productHandler.GetProducts)
	product.GET("/:id", productHandler.GetProduct)
	product.POST("", productHandler.CreateProduct)
	product.PUT("/:id", productHandler.UpdateProduct)
	product.DELETE("/:id", productHandler.DeleteProduct)
}
