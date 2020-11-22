package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"Authorization"},
	}))

	// Vueのbuild先のフォルダを指定
	e.Static("/", "../../../dist")

	e.GET("/hello", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":5000"))
}
