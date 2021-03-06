package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/utils"
	"log"
	"net/http"
)

/*

curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/Users/sam/docker-compose/sam/wwwroot/learn/go-gin/handbook/demo06/from/a.txt" \
  -F "upload[]=@/Users/sam/docker-compose/sam/wwwroot/learn/go-gin/handbook/demo06/from/b.txt" \
  -H "Content-Type: multipart/form-data"

 */

func main() {
	cpath:=utils.GetCurrentPath(1)
	fmt.Println(cpath)
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件到指定的路径
			dst:=cpath+"/to/"+file.Filename
			c.SaveUploadedFile(file,dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.Run(":8080")
}



