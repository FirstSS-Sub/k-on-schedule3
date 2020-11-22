package router

import (
	"github.com/FirstSS-Sub/k-on-schedule2/server/injector"
	"github.com/labstack/echo"
)

func UserRouting(e *echo.Echo) {
	// TODO REST API に沿ったpathの指定
	handler := injector.InjectUserHandler()

	e.Group("/user")

	e.GET("/", handler.FindById())

	e.POST("/", handler.Create())

	e.POST("/", handler.UpdateSchedule())

	e.POST("/", handler.ChangeName())

	e.POST("/", handler.Delete())
}
