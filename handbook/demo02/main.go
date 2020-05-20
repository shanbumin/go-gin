package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin的路由来自httprouter库。因此httprouter具有的功能，gin也具有，不过gin不支持路由正则表达式
//http://127.0.0.1:8080/user/sam
//http://127.0.0.1:8080/user/sam/eating
func main() {
	router := gin.Default()
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
    // 冒号:加上一个参数名组成路由参数。可以使用c.Params的方法读取其值。当然这个值是字符串string。
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	//除了:，gin还提供了*号处理参数，*号能匹配的规则就更多。
	//但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	//如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}