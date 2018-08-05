package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/shirakiya/sharqen/backend/conf"
	"github.com/shirakiya/sharqen/backend/database"
	"github.com/shirakiya/sharqen/backend/model"
)

func getCsvRows(filePath string) [][]string {
	fh, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	reader := csv.NewReader(fh)
	rows, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	return rows
}

func main() {
	config := conf.GetConfig()
	db := database.Init(config)
	defer db.Close()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	path := path.Join("seeds", "livequiz_dataset.csv")
	rows := getCsvRows(path)

	var recordCount int
	var notRecordCount int

	for _, row := range rows[1:] {
		var quiz model.Quiz
		var question string
		var choices []model.Choice
		var num int
		var correct int

		question = row[3]

		// skip creating record if same quiz is already exists.
		var existQuiz model.Quiz
		// search a quiz deleted softly to avoid creating duplicately
		tx.Unscoped().Where(&model.Quiz{Question: question}).First(&existQuiz)
		if existQuiz.ID != 0 {
			notRecordCount++
			continue
		}

		for _, choice_text := range row[4:8] {
			if choice_text == "" {
				continue
			}
			choices = append(choices, model.Choice{Text: choice_text})
		}
		num, _ = strconv.Atoi(row[2])
		correct, _ = strconv.Atoi(row[8])
		quiz = model.Quiz{
			ChannelName: row[1],
			Num:         num,
			Question:    row[3],
			Choices:     choices,
			Correct:     correct,
		}

		tx.Create(&quiz)
		recordCount++
	}

	if err := tx.Commit().Error; err == nil {
		fmt.Println("Finish to create seed data.")
		fmt.Println("new record:", recordCount)
		fmt.Println("skip record:", notRecordCount)
	} else {
		panic(err)
	}
}
