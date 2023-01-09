package main

import (
	"net/http"

	"github.com/keiji25/react-ts-tutorial/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	model.Connect()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.JSON(http.StatusNotFound, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}
		}
	}

	e.GET("/get/articles", model.GetArticles)
	e.POST("/create/articles", model.CreateArticles)
	e.POST("delete/:id", model.DeleteArticle)
	e.Logger.Fatal(e.Start(":8080"))
}
