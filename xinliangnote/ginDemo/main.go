package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/xinliangnote/ginDemo/config"
	"go-gin/xinliangnote/ginDemo/middleware"
	"go-gin/xinliangnote/ginDemo/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	//增肌一个日志中间件
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}