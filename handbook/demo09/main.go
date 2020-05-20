package main

import (
	"github.com/gin-gonic/gin"
)

/*
使用  r := gin.New()  代替   默认启动方式，包含 Logger、Recovery 中间件  r := gin.Default()
*/



func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()
	r.Run(":8080")
}