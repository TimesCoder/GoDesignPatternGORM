package routes

import (
	"design_pattern/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBookByID)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)
}
