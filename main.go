package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func main()  {
	app :=dotweb.New()
	app.SetDevelopmentMode()
	app.SetEnabledLog(true)
	app.UseRequestLog()

	app.HttpServer.GET("/index", func(context dotweb.Context) error {

		return context.WriteString("welcome to my first web!")

	})



	fmt.Println("dotweb.StartServer begin")
	err := app.StartServer(80)

	fmt.Println("dotweb.StartServer error => ", err)
}