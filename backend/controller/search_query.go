package controller

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/shirakiya/sharqen/backend/model"
)

type (
	SearchQueryRequestParams struct {
		QuizID      int    `json:"quiz_id"`
		SearchQuery string `json:"search_query"`
	}

	OKResponse struct {
		Status string `json:"status"`
	}
)

func RoutePostSearchQuery(c echo.Context) error {
	params := new(SearchQueryRequestParams)
	if err := c.Bind(params); err != nil {
		fmt.Println(err)
		return err
	}

	tx := c.Get("tx").(*gorm.DB)

	quiz := model.Quiz{}
	tx.First(&quiz, params.QuizID)

	quiz.SearchQuery = params.SearchQuery
	tx.Save(&quiz)

	return c.JSON(http.StatusOK, OKResponse{Status: "OK"})
}
