package  main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context){
	c.String(http.StatusOK, "getting\n")
}
func posting(c *gin.Context){
	c.String(http.StatusOK, "posting\n")
}

func putting(c *gin.Context){
	c.String(http.StatusOK, "putting\n")
}
func deleting(c *gin.Context){
	c.String(http.StatusOK, "deleting\n")
}
func patching(c *gin.Context){
	c.String(http.StatusOK, "patching\n")
}
func head(c *gin.Context){
	c.String(http.StatusOK, "head\n")
}
func options(c *gin.Context){
	c.String(http.StatusOK, "options\n")
}

/*
使用gin的Default方法创建一个路由handler。然后通过HTTP方法绑定路由规则和路由函数。
不同于net/http库的路由函数，gin进行了封装，把request和response都封装到gin.Context的上下文环境。(w http.ResponseWriter, r *http.Request)
最后是启动路由的Run方法监听端口。麻雀虽小，五脏俱全。当然，除了GET方法，gin也支持POST,PUT,DELETE,OPTION等常用的restful方法。
*/
func main() {
// Disable Console Color
// gin.DisableConsoleColor()
// 使用默认中间件创建一个gin引擎 ,默认是 Logger and  Recovery  (crash-free) 中间件
router := gin.Default()

router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)
// 默认启动的是 8080端口，也可以自己定义启动端口 router.Run(":3000")
router.Run()

}