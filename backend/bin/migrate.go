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

	db.AutoMigrate(
		&model.Quiz{},
		&model.Choice{},
	)
	db.Model(&model.Choice{}).AddForeignKey("quiz_id", "quizzes(id)", "CASCADE", "CASCADE")

	fmt.Println("Finish migration.")
}
