package routers

import (
	"dtcms/routers/v1"
	"github.com/devfeel/dotweb"
)

func InitRouter() *dotweb.DotWeb  {
	app :=dotweb.New()

	app.HttpServer.GET("/index", func(context dotweb.Context) error {

		return context.WriteString("welcome to my first web!")
	})

	g := app.HttpServer.Group("/v1")

	g.GET("/menulist",v1.QueryMenuLists)

	return app
}