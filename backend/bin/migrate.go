package main

import (
	"fmt"

	"github.com/shirakiya/sharqen/backend/conf"
	"github.com/shirakiya/sharqen/backend/database"
	"github.com/shirakiya/sharqen/backend/model"
)

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

	tx.AutoMigrate(
		&model.Quiz{},
		&model.Choice{},
	)
	tx.Model(&model.Choice{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")

	if err := tx.Commit().Error; err == nil {
		fmt.Println("Finish migration.")
	} else {
		panic(err)
	}
}
