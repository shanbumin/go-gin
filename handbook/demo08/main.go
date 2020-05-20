package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func loginEndpoint(c *gin.Context){
	c.String(http.StatusOK, "loginEndpoint\n")
}
func submitEndpoint(c *gin.Context){
	c.String(http.StatusOK, "submitEndpoint\n")
}

func readEndpoint(c *gin.Context){
	c.String(http.StatusOK, "readEndpoint\n")
}


//curl -v  -X POST   http://127.0.0.1:8080/v1/login
//curl -v  -X POST   http://127.0.0.1:8080/v2/login
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}