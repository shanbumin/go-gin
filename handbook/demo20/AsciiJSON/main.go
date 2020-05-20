package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用AsciiJSON将使特殊字符编码
//http://localhost:8080/someJSON  ===>浏览器会自动转换的，所以看得不明显

//curl -v http://localhost:8080/someJSON


func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 将输出: {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		//c.JSON(http.StatusOK, data)
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
