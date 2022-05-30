package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	port := "https://108.136.160.138:443"
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		u := new(Message)
		u.Message = "Hello World!"
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(port))
}
