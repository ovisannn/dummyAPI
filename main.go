package main

import (
	book "dummyAPI/model"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "check_echo", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}

func main() {
	e := echo.New()
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bookGroup := e.Group("/book")
	{
		bookGroup.POST("", book.Create)
		bookGroup.PUT("/:id", book.Update)
		bookGroup.DELETE("/:id", book.Delete)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
