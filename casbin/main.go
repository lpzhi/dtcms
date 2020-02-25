package main

import (
	"fmt"
	"github.com/casbin"
	"github.com/gorm-adapter"
	"github.com/devfeel/dotweb"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_"dtcms/models"
)

// 数据库连接及角色规则的初始化
func connect() {
	dsn := "root:0987ABc123.@(127.0.0.1:3306)/dotcms?charset=utf8&parseTime=True&loc=Local&allowNativePasswords=true"
	var err error
	or, err := gorm.Open("mysql", dsn)
	defer or.Close()

	if err != nil {
		fmt.Println("connect DB error",err.Error())
		return
		//panic(err)
	}

	// 将数据库连接同步给插件， 插件用来操作数据库
	po,_ := gormadapter.NewAdapterByDB(or)
	// 这里也可以使用原生字符串方式
	//
	e,err := casbin.NewEnforcer("/Users/lupeng/golang/src/dtcms/casbin/rbac_model.conf", po)

	if err!=nil {
		fmt.Println(err.Error())

	}
	// 开启权限认证日志
	e.EnableLog(true)
	// 加载数据库中的策略
	err = e.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error",err.Error())

		//panic(err)
	}

	fmt.Println("aa")
	// 创建一个角色,并赋于权限
	// admin 这个角色可以访问GET 方式访问 /api/v2/ping
	res,_ := e.AddPolicy("admin","/api/v2/ping","GET")
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	// 将 test 用户加入一个角色中
	e.AddRoleForUser("test","root")
	e.AddRoleForUser("tom","admin")
	// 请看规则中如果用户名为 root 则不受限制
}

func main()  {

	app :=dotweb.New()
	app.SetDevelopmentMode()
	app.SetEnabledLog(true)
	app.UseRequestLog()

	app.HttpServer.GET("/index", func(context dotweb.Context) error {
		connect()
		fmt.Println("connect success")
		return context.WriteString("welcome to my first web!")

	})



	fmt.Println("dotweb.StartServer begin")
	err := app.StartServer(9091)

	fmt.Println("dotweb.StartServer error => ", err)
}
