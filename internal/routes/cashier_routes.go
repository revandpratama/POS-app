package routes

import (
	"point-of-sales-app/adapter"
	"point-of-sales-app/internal/handlers"
	"point-of-sales-app/internal/repositories"
	"point-of-sales-app/internal/services"

	"github.com/labstack/echo/v4"
)

func initCashierHandler() handlers.CashierHandler {
	repo := repositories.NewCashierRepository(adapter.DB)
	service := services.NewCashierService(repo)
	return handlers.NewCashierHandler(service)
}

func InitCashierRoutes(e *echo.Group) {
	cashierHandler := initCashierHandler()

	cashier := e.Group("/cashier")

	cashier.GET("/transactions", cashierHandler.GetTransactionList)
	cashier.POST("/transactions", cashierHandler.CreateTransaction)
	cashier.PUT("/transactions/:id", cashierHandler.UpdateStock)
}
