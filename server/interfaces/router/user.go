package router

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/injector"
	"github.com/labstack/echo"
)

func UserRouting(e *echo.Echo) {
	// TODO REST API に沿ったpathの指定
	handler := injector.InjectUserHandler()

	g := e.Group("/user")

	g.GET("/", handler.FindByUid(), jwtAuth())

	g.POST("/", handler.Create())

	g.PUT("/schedule", handler.UpdateSchedule(), jwtAuth())

	g.PUT("/", handler.ChangeName(), jwtAuth())

	g.DELETE("/", handler.Delete(), jwtAuth())
}
