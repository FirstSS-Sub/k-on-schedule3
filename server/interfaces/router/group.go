package router

import (
	"github.com/FirstSS-Sub/k-on-schedule3/server/injector"
	"github.com/labstack/echo"
)

func GroupRouting(e *echo.Echo) {
	// TODO REST API に沿ったpathの指定
	handler := injector.InjectGroupHandler()

	g := e.Group("/group")

	g.GET("/", handler.FindById(), jwtAuth())

	g.POST("/", handler.Create(), jwtAuth())

	g.PUT("/", handler.ChangeName(), jwtAuth())

	g.POST("/member", handler.AddUser(), jwtAuth())

	g.DELETE("/member", handler.Leave(), jwtAuth())

	g.DELETE("/", handler.Delete(), jwtAuth())
}
