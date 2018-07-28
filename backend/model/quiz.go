package model

import (
	"github.com/jinzhu/gorm"
)

// type Channel struct {
//     gorm.Model
//     name string `sql:"type:varchar(255)"`
// }

type Quiz struct {
	gorm.Model
	// ChannelId int
	// Channel Channel
	ChannelName string `sql:"type:varchar(255);not null;"`
	Num         int    `sql:"not null;default:0;"`
	Question    string `sql:"type:text; not null;"`
	Choices     []Choice
	Correct     int    `sql:"not null;"`
	SearchQuery string `sql:"type:text;default:null;"`
}

type Choice struct {
	gorm.Model
	QuizId int `sql:"not null;index;"`
	Quiz   Quiz
	Text   string `sql:"type:varchar(255);not null;"`
}
