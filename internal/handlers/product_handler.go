package handlers

import (
	"net/http"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	GetProducts(c echo.Context) error
	GetProduct(c echo.Context) error
	CreateProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}
type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}

func (h *productHandler) GetProducts(c echo.Context) error {

	products, err := h.productService.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    products,
	})
}

func (h *productHandler) GetProduct(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	product, err := h.productService.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    product,
	})
}

func (h *productHandler) CreateProduct(c echo.Context) error {

	var req dto.ProductRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := h.productService.CreateProduct(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (h *productHandler) UpdateProduct(c echo.Context) error {

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

	if err := h.productService.UpdateProduct(id, &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update product",
	})
}

func (h *productHandler) DeleteProduct(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := h.productService.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete product",
	})
}
