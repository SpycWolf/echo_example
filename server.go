package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		jsonMap := map[string]string{
			"name": "taro",
			"hoge": "moge",
		}
		return c.JSON(http.StatusOK, jsonMap)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
