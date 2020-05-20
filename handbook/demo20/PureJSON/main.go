package main

import "github.com/gin-gonic/gin"

//通常情况下，JSON会将特殊的HTML字符替换为对应的unicode字符，比如<替换为\u003c，如果想原样输出html，则使用PureJSON，这个特性在Go 1.6及以下版本中无法使用。


func main() {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}



