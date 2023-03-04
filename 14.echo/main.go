package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルート
	e.GET("/", hiyoko)
	e.Logger.Fatal(e.Start(":8000"))
}

func hiyoko(c echo.Context) error {
	return c.String(http.StatusOK, "hiyoko")
}
