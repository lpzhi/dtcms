package main

import (
	"dtcms/config"
	"dtcms/global"
	"dtcms/models"
	"dtcms/routers"
	"dtcms/util/file"
	"flag"
	"fmt"
	"os"
	"strconv"
)


var (
	configFile string
	configPath string
	RunEnv     string
)

const (
	RunEnv_Flag       = "RunEnv"
	RunEnv_Develop    = "develop"
	RunEnv_Test       = "test"
	RunEnv_Production = "production"
)

func parseFlag() {
	RunEnv = os.Getenv(RunEnv_Flag)
	if RunEnv == "" {
		RunEnv = RunEnv_Develop
	}

	configPath = _file.GetCurrentDirectory() + "/conf/" + RunEnv
	//load app config
	flag.StringVar(&configFile, "config", "", "配置文件路径")
	if configFile == "" {
		configFile = configPath + "/app.conf"
	}

}

func main()  {

	app := routers.InitRouter()
	app.SetDevelopmentMode()
	app.SetEnabledLog(true)
	app.UseRequestLog()


	parseFlag()
	//加载全局xml配置文件
	config.InitConfig(configFile)


	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}

	fmt.Println("dotweb.StartServer begin1",configPath)
	fmt.Println(_file.GetCurrentPath())
	//全局初始化

	er := global.Init(configPath)
	if er != nil {
		panic(er)
	}


	//fmt.Println()
	//fmt.Println(models.GetDb())
	//db connect
	models.ConnectDb()
	err := app.StartServer(8080)

	fmt.Println("dotweb.StartServer error => ", err)
}