package router

import (
	"github.com/labstack/echo"

	"github.com/shirakiya/sharqen/backend/controller"
)

func Init(e *echo.Echo) *echo.Echo {
	e.GET("/quiz/next", controller.RouteGetNextQuiz)
	e.GET("/search-result", controller.RouteGetSearchResult)
	e.POST("/search-query", controller.RoutePostSearchQuery)

	return e
}
