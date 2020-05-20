package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//http://localhost:8080/someJSON

//使用SecureJSON可以防止json劫持，如果返回的数据是数组，则会默认在返回值前加上"while(1)"

func main() {
	r := gin.Default()

	// 可以自定义返回的json数据前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

