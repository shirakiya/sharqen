package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/shirakiya/sharqen/backend/model"
	"github.com/shirakiya/sharqen/backend/utils"
)

type (
	FormattedQuiz struct {
		Id          uint     `json:"id"`
		ChannelName string   `json:"channel_name"`
		Num         int      `json:"num"`
		Question    string   `json:"question"`
		Choices     []string `json:"choices"`
		Correct     int      `json:"correct"`
	}

	QuizResponse struct {
		Quiz  FormattedQuiz `json:"quiz"`
		Query string        `json:"query"`
	}
)

func FormatToResponse(quiz model.Quiz, query string) QuizResponse {
	var choice_texts []string
	for _, choice := range quiz.Choices {
		choice_texts = append(choice_texts, choice.Text)
	}

	formattedQuiz := FormattedQuiz{
		Id:          quiz.ID,
		ChannelName: quiz.ChannelName,
		Num:         quiz.Num,
		Question:    quiz.Question,
		Choices:     choice_texts,
		Correct:     quiz.Correct,
	}

	return QuizResponse{
		Quiz:  formattedQuiz,
		Query: query,
	}
}

func GetNextQuiz(tx *gorm.DB) model.Quiz {
	quiz := model.Quiz{}

	tx.Where("search_query IS NULL").First(&quiz)
	tx.Model(&quiz).Related(&quiz.Choices)

	return quiz
}

func RouteGetQuiz(c echo.Context) error {
	tx := c.Get("tx").(*gorm.DB)

	quiz := GetNextQuiz(tx)

	var choices []string
	for _, choice := range quiz.Choices {
		choices = append(choices, choice.Text)
	}
	customQueryResponse := utils.GetCustomQuery(quiz.Question, choices)

	response := FormatToResponse(quiz, customQueryResponse.CustomQuery)

	return c.JSON(http.StatusOK, response)
}
