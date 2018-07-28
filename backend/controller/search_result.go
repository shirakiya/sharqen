package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/shirakiya/sharqen/backend/utils"
)

type SearchResultRequestParams struct {
	Question string   `query:"question"`
	Choices  []string `query:"choices[]"`
}

func RouteGetSearchResult(c echo.Context) error {
	params := new(SearchResultRequestParams)
	if err := c.Bind(params); err != nil {
		fmt.Println(err)
		return err
	}

	solveResultResponse := utils.GetSolveResult(params.Question, params.Choices)

	return c.JSON(http.StatusOK, solveResultResponse)
}
