package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


//发送数据给服务端，并不是post方法才行，put方法一样也可以。同时querystring和body也不是分开的，两个同时发送也可以：
//curl -v  -X POST  -d "name=manu&message=this_is_great"    http://127.0.0.1:8080/post?id=1234&page=1
func main() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
        //id: 1234; page: 0; name: manu; message: this_is_great
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}
//上面的例子，展示了同时使用查询字串和body参数发送数据给服务器。




