package main

import (
	"Book_Management/pkg/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.RegisterBookStoreRoutes(e)
	e.Logger.Fatal(e.Start(":9010"))
}
