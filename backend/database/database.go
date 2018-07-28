package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/shirakiya/sharqen/backend/conf"
)

func createConnectionStatement(c conf.Config) string {
	// ex. "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)
}

func Init(c conf.Config) *gorm.DB {
	db, err := gorm.Open("mysql", createConnectionStatement(c))
	if err != nil {
		panic(err)
	}

	return db
}
