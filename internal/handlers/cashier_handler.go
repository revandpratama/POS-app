package handlers

import (
	"net/http"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CashierHandler interface {
	CreateTransaction(c echo.Context) error
	GetTransactionList(c echo.Context) error
	UpdateStock(c echo.Context) error
}

type cashierHandler struct {
	service services.CashierService
}

func NewCashierHandler(service services.CashierService) CashierHandler {
	return &cashierHandler{service: service}
}

func (h *cashierHandler) CreateTransaction(c echo.Context) error {

	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := h.service.CreateTransaction(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success creating transaction",
	})
}

func (h *cashierHandler) GetTransactionList(c echo.Context) error {

	transactions, err := h.service.GetTransactionList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    transactions,
	})
}

func (h *cashierHandler) UpdateStock(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := h.service.UpdateStock(id, req.Quantity); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update product stock",
	})
}
