package basicFramework

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var users []User

func GetUsersController(e echo.Context) error {
	return e.JSON(http.StatusOK, Response{
		Message: "Success",
		Data:    users,
	})
}

func GetUserController(e echo.Context) error {
	id := e.Param("id")
	for _, user := range users {
		if id == strconv.Itoa(user.Id) {
			return e.JSON(http.StatusOK, Response{
				Message: "Success",
				Data:    user,
			})
		}
	}
	return e.JSON(http.StatusNotFound, Response{
		Message: "User Not Found",
		Data:    nil,
	})
}

func CreateUserController(e echo.Context) error {
	// binding data
	user := User{}
	e.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return e.JSON(http.StatusCreated, Response{
		Message: "Success",
		Data:    user,
	})
}

func UpdateUserController(e echo.Context) error {
	id := e.Param("id")

	updatedUser := User{}
	e.Bind(&updatedUser)

	index := -1
	for i, user := range users {
		if id == strconv.Itoa(user.Id) {
			index = i
		}
	}

	if index != -1 {
		updatedUser.Id, _ = strconv.Atoi(id)
		users[index] = updatedUser
		return e.JSON(http.StatusOK, Response{
			Message: "Success",
			Data:    users[index],
		})
	}

	return e.JSON(http.StatusNotFound, Response{
		Message: "User Not Found",
		Data:    nil,
	})
}

func DeleteUserController(e echo.Context) error {
	id := e.Param("id")

	var updatedUsers []User
	for _, user := range users {
		if id == strconv.Itoa(user.Id) {
			continue
		} else {
			updatedUsers = append(updatedUsers, user)
			users = updatedUsers
		}
	}

	return e.JSON(http.StatusOK, Response{
		Message: "Success",
		Data:    nil,
	})
}

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	return e
}
