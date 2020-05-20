package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//在中间件或处理程序中启动新的Goroutines时，你不应该使用其中的原始上下文，你必须使用只读副本（c.Copy()）
//http://localhost:8080/long_sync
//http://localhost:8080/long_async
func main() {
	r := gin.Default()

	//异步
	r.GET("/long_async", func(c *gin.Context) {
		// 创建要在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// 这里使用你创建的副本
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})
	//同步
	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// 这里没有使用goroutine，所以不用使用副本
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}


