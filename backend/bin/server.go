package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/shirakiya/sharqen/backend/conf"
	"github.com/shirakiya/sharqen/backend/database"
	myMiddleware "github.com/shirakiya/sharqen/backend/middleware"
	"github.com/shirakiya/sharqen/backend/router"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	config := conf.GetConfig()

	db := database.Init(config)
	db.LogMode(true)
	defer db.Close()

	e.Use(myMiddleware.TransactionHandler(db))

	e = router.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
