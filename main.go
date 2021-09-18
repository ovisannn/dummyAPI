package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hellocontroller(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!")
}

func main() {
	e := echo.New()

	e.GET("/", Hellocontroller)

	e.Start(":4000")
}
