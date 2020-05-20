package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//访问静态文件需要先设置路径
//原生的
// handle static assets
//mux := http.NewServeMux()
//files := http.FileServer(http.Dir(config.Static))
//mux.Handle("/static/", http.StripPrefix("/static/", files))


//@todo http://127.0.0.1:8080/assets/a.html
//@todo http://127.0.0.1:8080/more_static/b.html
//@todo http://127.0.0.1:8080/favicon.ico

//StaticFile 是加载单个文件，而StaticFS 是加载一个完整的目录资源：
func main() {
	//router.Static("/static/", "./static")
	router := gin.Default()

	fmt.Println(http.Dir("my_file_system"))
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.Run(":8080")
}

