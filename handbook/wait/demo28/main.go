package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//直接像这样使用http.ListenAndServe()
	//router := gin.Default()
	//http.ListenAndServe(":8080", router)

	//router就是一个自定义的多路复用器
	router := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}



