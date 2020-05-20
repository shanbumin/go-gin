package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//快速体验
	//router := gin.Default()
	//router.LoadHTMLGlob("templates/*")
	////router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	//router.Run(":8080")

	//在不同目录中使用具有相同名称的模板
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	router.Run(":8080")
}




