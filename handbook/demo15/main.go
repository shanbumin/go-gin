package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}
//http://localhost:8085/testing?name=shanbumin&address=jingsu
func startPage(c *gin.Context) {
	var person Person
	//ShouldBindQuery 函数只绑定Get参数，不绑定post数据
	if c.ShouldBindQuery(&person) == nil {
		fmt.Println(person)
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}