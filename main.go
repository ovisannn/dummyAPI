package main

import (
	"net/http"

	"dummyAPI/connection"

	"github.com/labstack/echo/v4"
)

func Hellocontroller(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}

func main() {

	urldb := "mongodb://localhost:27017"

	client, ctx, cancel, err := connection.Connect(urldb)
	if err != nil {
		panic(err)
	}

	connection.Ping(client, ctx)

	e := echo.New()

	e.GET("/", Hellocontroller)

	e.Start(":4000")
	defer connection.Close(client, ctx, cancel)
}
