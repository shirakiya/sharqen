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

func main() {
	config := conf.GetConfig()
	db := database.Init(config)
	defer db.Close()

	fh, err := os.Open(path.Join("seeds", "livequiz_dataset.csv"))
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	reader := csv.NewReader(fh)
	rows, err := reader.ReadAll()

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, row := range rows[1:] {
		var quiz model.Quiz
		var choices []model.Choice
		var num int
		var correct int

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
	}
	err = tx.Commit().Error

	if err == nil {
		fmt.Println("Finish to create seed data.")
	} else {
		fmt.Println(err)
	}
}
