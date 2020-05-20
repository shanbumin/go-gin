package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

//https://github.com/fvbock/endless/tree/master/examples
//想要优雅地重启或停止你的Web服务器，使用下面的方法
//我们可以使用fvbock/endless来替换默认的ListenAndServe
//https://blog.csdn.net/tomatomas/article/details/94839857

//https://blog.csdn.net/black_OX/article/details/77869500?locationNum=9&fps=1  [热重启golang服务器]
//https://www.cnblogs.com/CraryPrimitiveMan/p/8560839.html [Golang学习--平滑重启]

func main()  {

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "just a test")
	})
	s := endless.NewServer(":8199", r)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("server err: %v", err)
	}


}