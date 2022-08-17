package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.GET("/", hello)
	server.Logger.Fatal(server.Start(":8080"))
}
func hello(c echo.Context) error {
	toJSON := &ToJSON{"Hello world"}
	return c.JSON(http.StatusOK, toJSON)
}

type ToJSON struct {
	Text string `json:"Text"`
}
