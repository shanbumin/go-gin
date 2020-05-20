package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required" time_format:"2006-01-02"`
}

//curl "localhost:8085/bookable?check_in=2019-04-16&check_out=2019-04-17"
//curl "localhost:8085/bookable?check_in=2018-03-08"
func main() {
	route := gin.Default()
	//路由
	route.GET("/bookable", getBookable)
	route.Run(":8085")
}


//处理函数
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!","in":b.CheckIn,"out":b.CheckOut})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}


