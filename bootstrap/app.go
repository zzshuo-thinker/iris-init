package bootstrap

import (
	"iris-init/config"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"github.com/kataras/iris"
	"strconv"
	"github.com/kataras/iris/middleware/logger"
			"iris-init/routers"
	"log"
	"iris-init/common"
)

type Application struct {
	Application *iris.Application
	Debug       bool
}

func App() *Application {
	application := &Application{Application: iris.New()}
	application.initEnv()
	application.isDebug()
	application.initDB()
	application.initRouter()
	return application
}

// 初始化数据库
func (app *Application) initDB()  {
	config.DatabaseConn()
}

// 初始化配置文件
func (app *Application) initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[-] Error loading .env file")
	}
	fmt.Println("[::] APP Running on Port " + os.Getenv("API_PORT"))
}

// 初始化debug
func (app *Application) isDebug()  {
	debug := os.Getenv("DEBUG")
	if debug != "" {
		app.Debug, _ = strconv.ParseBool(debug)
	}
}

// 初始化日志
func (app *Application) initLogger() {
	requestLogger := logger.New(logger.Config{
		Status: 				true,
		IP:						true,
		Method:					true,
		Path: 					true,
		Query:					true,
		MessageContextKeys: 	[]string{"logger message"},
		MessageHeaderKeys: 		[]string{"User-Agent"},
	})
	app.Application.Use(requestLogger)
}

// 初始化路由
func (app *Application) initRouter() {
	app.initLogger()
	app.Application.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(common.ErrorClientParams)
	})

	app.Application.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.JSON(common.ErrorUnKnow)
	})
	routers.Router(app.Application)
}


func (app *Application) runner() iris.Runner {
	return iris.Addr(":8888")
}
// run
func (app *Application) Run()  {
	err := app.Application.Run(app.runner())
	if err != nil {
		log.Fatalf("Init application run err : %v", err)
	}
}