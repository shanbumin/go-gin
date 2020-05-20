package main

import "github.com/gin-gonic/gin"

//http的报文体传输数据就比query string稍微复杂一点，常见的格式就有四种。
//例如application/json，application/x-www-form-urlencoded, application/xml和multipart/form-data。后面一个主要用于图片上传。
//json格式的很好理解，urlencode其实也不难，无非就是把query string的内容，放到了body体里，同样也需要urlencode。
//默认情况下，c.PostFROM解析的是x-www-form-urlencoded或from-data的参数。
//与get处理query参数一样，post方法也提供了处理默认参数的情况。同理，如果参数不存在，将会得到空字串。

//curl -v  -X POST  -d "message=hello&nick=shit&age=18"  http://127.0.0.1:8080/form_post
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}


//前面我们使用c.String返回响应，顾名思义则返回string类型。content-type是plain或者text。
//调用c.JSON则返回json数据。其中gin.H封装了生成json的方式，是一个强大的工具。
//使用golang可以像动态语言一样写字面量的json，对于嵌套json的实现，嵌套gin.H即可。



