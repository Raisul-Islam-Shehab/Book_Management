package routes

import(
	"github.com/labstack/echo/v4"
	"Book_Management/pkg/controllers"
)

func RegisterBookStoreRoutes(e *echo.Echo){
	e.GET("/book/", controllers.GetBook)
	e.POST("/book/", controllers.CreateBook)
	e.GET("/book/:bookId", controllers.GetBookById)
	e.PUT("/book/:bookId", controllers.UpdateBookById)
	e.DELETE("/book/:bookId", controllers.DeleteBookById)
}


