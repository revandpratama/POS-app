package handlers

import (
	"net/http"
	"point-of-sales-app/internal/dto"
	"point-of-sales-app/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) GetUsers(c echo.Context) error {

	users, err := u.userService.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    users,
	})
}

func (u *userHandler) GetUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	user, err := u.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"data":    user,
	})
}

func (u *userHandler) UpdateUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var req dto.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := u.userService.UpdateUser(id, &req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update user",
	})
}

func (u *userHandler) DeleteUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	if err := u.userService.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete user",
	})
}
