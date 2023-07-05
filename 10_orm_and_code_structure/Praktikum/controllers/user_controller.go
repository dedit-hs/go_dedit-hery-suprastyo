package controllers

import (
	"net/http"

	"github.com/dedit-hs/go_dedit-hery-suprastyo/10_orm_and_code_structure/Praktikum/config"
	"github.com/dedit-hs/go_dedit-hery-suprastyo/10_orm_and_code_structure/Praktikum/models"
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
