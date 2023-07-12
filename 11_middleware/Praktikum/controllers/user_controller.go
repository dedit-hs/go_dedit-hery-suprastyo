package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_dedit-hery-suprastyo/11_middleware/Praktikum/config"
	"github.com/dedit-hs/go_dedit-hery-suprastyo/11_middleware/Praktikum/middlewares"
	"github.com/dedit-hs/go_dedit-hery-suprastyo/11_middleware/Praktikum/models"
	"github.com/labstack/echo/v4"
)

func GetUsersController(e echo.Context) error {
	var users []models.User
	err := config.DB.Find(&users)

	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    users,
	})
}

func GetUserController(e echo.Context) error {
	id := e.Param("id")
	var user models.User

	err := config.DB.First(&user, id)

	if err.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    user,
	})
}

func AddUserController(e echo.Context) error {
	var user models.User
	e.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusCreated, models.Response{
		Message: "Success",
		Data:    user,
	})
}

func DeleteUserController(e echo.Context) error {
	id := e.Param("id")
	var user models.User
	found := config.DB.First(&user, id)

	if found.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	deleteUser := config.DB.Delete(&user)
	if deleteUser.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    nil,
	})
}

func UpdateUserController(e echo.Context) error {
	id := e.Param("id")
	var user models.User

	found := config.DB.First(&user, id)

	if found.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	e.Bind(&user)
	updatedUser := config.DB.Save(&user)

	if updatedUser.Error != nil {
		return e.JSON(http.StatusInternalServerError, models.Response{
			Message: "Internal Server Error",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, models.Response{
		Message: "Success",
		Data:    user,
	})
}

func LoginUserController(e echo.Context) error {
	var user models.User
	e.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login failed",
			"error":   err.Error(),
		})
	}

	token, err := middlewares.CreateToken(user.Id, user.Name)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login failed",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{Id: user.Id, Name: user.Name, Email: user.Email, Token: token}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"user":    userResponse,
	})
}
