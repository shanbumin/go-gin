package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


//web提供的服务通常是client和server的交互。其中客户端向服务器发送请求，除了路由参数，其他的参数无非两种，查询字符串query string和报文体body参数。
//所谓query string，即路由，用?以后连接的key1=value1&key2=value2的形式的参数。当然这个key-value是经过urlencode编码。
//使用c.DefaultQuery方法读取参数，其中当参数不存在的时候，提供一个默认值。使用Query方法读取正常参数，当参数不存在的时候，返回空字串：
//之所以使用中文，是为了说明urlencode。注意，当firstname为空字串的时候，并不会使用默认的Guest值，空值也是值，DefaultQuery只作用于key不存在的时候，提供默认值。

//@link http://127.0.0.1:8080/welcome?firstname=Jane&lastname=Doe
func main() {
	router := gin.Default()
	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}