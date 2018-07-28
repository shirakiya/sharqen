package middleware

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func TransactionHandler(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			tx := db.Begin()
			c.Set("tx", tx)

			if err := next(c); err != nil {
				tx.Rollback()
				fmt.Println("Transction Rollback: ", err)
				return err
			}
			tx.Commit()

			return nil
		})
	}
}
