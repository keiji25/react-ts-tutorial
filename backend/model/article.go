package model

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Article struct {
		ID        uint64 `gorm:"primary_key"`
		Content   string `json:"content" gorm:"type:varchar(255);not null"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
)

func GetArticles(c echo.Context) error {
	var articles []Article
	DB.Find(&articles)

	return c.JSON(http.StatusOK, articles)
}

func CreateArticles(c echo.Context) error {
	r := new(Article)
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	DB.Create(r)
	return c.JSON(http.StatusOK, nil)
}

func DeleteArticle(c echo.Context) error {
	id := c.Param("id")
	r := new(Article)
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	r.ID = uid
	DB.Delete(r)
	return c.JSON(http.StatusOK, nil)
}
