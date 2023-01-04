package main

import (
	"github.com/keiji25/react-ts-tutorial/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	model.Connect()
	e.Logger.Fatal(e.Start(":8080"))
}
