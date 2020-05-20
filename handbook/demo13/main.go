package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// 绑定为json
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	//Password string `form:"password" json:"password" xml:"password" binding:"-"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//curl -v -X POST   http://localhost:8080/loginJSON  -H 'content-type: application/json'  -d '{ "user": "manu" }'
//curl -v -X POST   http://localhost:8080/loginJSON  -H 'content-type: application/json'  -d '{ "user": "manu" ,"password":123}'
//curl -v -X POST   http://localhost:8080/loginJSON  -H 'content-type: application/json'  -d '{ "user": "manu","password":"123"}'
//@link https://www.kancloud.cn/shuangdeyu/gin_book/949426
//@link curl  -v  -X POST   http://localhost:8080/loginForm   -d  'user=manu&password=123'

func main() {
	router := gin.Default()

	// ①Example for binding JSON
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		//ShouldBindJSON:将json格式的body内容转结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	//-----------------------------------------------------------------
	//②Example for binding XML
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		//ShouldBindJSON:将xml格式的body内容转结构体
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%+v",xml)
		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//③Example for binding a HTML form
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}



