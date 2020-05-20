package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/xinliangnote/ginDemo2/config"
	"go-gin/xinliangnote/ginDemo2/router"
)


func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.New()
	router.InitRouter(engine) // 设置路由
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
