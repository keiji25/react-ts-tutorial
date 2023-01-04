package main

import (
	"fmt"

	"github.com/keiji25/react-ts-tutorial/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	fmt.Println(model.DB)
	e.Logger.Fatal(e.Start(":8080"))
}
