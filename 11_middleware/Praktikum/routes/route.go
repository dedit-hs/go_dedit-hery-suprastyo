package routes

import (
	"github.com/dedit-hs/go_dedit-hery-suprastyo/11_middleware/Praktikum/constants"
	"github.com/dedit-hs/go_dedit-hery-suprastyo/11_middleware/Praktikum/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	eJWT := e.Group("")
	eJWT.Use(echojwt.JWT([]byte(constants.JWT_SECRET)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.AddUserController)
	e.POST("/login", controllers.LoginUserController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserController)
	eJWT.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	eJWT.POST("/books", controllers.AddBookController)
	eJWT.DELETE("/books/:id", controllers.DeleteBookController)
	eJWT.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
