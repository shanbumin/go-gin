package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用JSONP可以跨域传输，如果参数中存在回调参数，那么返回的参数将是回调函数的形式
//http://localhost:8080/JSONP?callback=call


func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
			"foo2": "bar2",
		}

		// 访问 http://localhost:8080/JSONP?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}



