package main

import (
	"github.com/gin-gonic/gin"
)

//发布HTTP重定向很容易，支持内部和外部链接
func main() {

	r := gin.Default()
	//外部跳转
	//r.GET("/test", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	//})


	//Gin路由重定向，使用如下的HandleContext
	r.GET("/test-inner", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})


	r.Run(":8080")


	
}
