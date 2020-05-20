package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func MyBenchLogger(c *gin.Context){
	c.String(http.StatusOK, "MyBenchLogger\n")
}

func benchEndpoint(c *gin.Context){
	c.String(http.StatusOK, "benchEndpoint\n")
}


func AuthRequired(c *gin.Context){
	//fmt.Println("AuthRequired\n")
	c.String(http.StatusOK, "AuthRequired\n")

}


func loginEndpoint(c *gin.Context){
	c.String(http.StatusOK, "loginEndpoint\n")
}

func submitEndpoint(c *gin.Context){
	c.String(http.StatusOK, "submitEndpoint\n")
}

func readEndpoint(c *gin.Context){
	c.String(http.StatusOK, "submitEndpoint\n")
}



func analyticsEndpoint(c *gin.Context){
	c.String(http.StatusOK, "analyticsEndpoint\n")
}




//http://127.0.0.1:8080/benchmark
//curl -v -X POST    http://127.0.0.1:8080/login
//http://127.0.0.1:8080/testing/analytics



func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()
	//①全局中间件
	r.Use(gin.Logger())   // 使用 Logger 中间件
	r.Use(gin.Recovery())  // 使用 Recovery 中间件

	//②路由添加中间件，可以添加任意多个
	r.GET("/benchmark", MyBenchLogger, benchEndpoint)
	//③路由组中添加中间件
	authorized := r.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group   嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}


	r.Run(":8080")
}