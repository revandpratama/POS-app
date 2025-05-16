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

	e.GET("/cashier/transactions", cashierHandler.GetTransactionList)
	e.POST("/cashier/transactions", cashierHandler.CreateTransaction)
	e.PUT("/cashier/transactions/:id", cashierHandler.UpdateStock)
}
