package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-gin/xinliangnote/ginDemo2/middleware/logger"
	"go-gin/xinliangnote/ginDemo2/middleware/sign"
	v1 "go-gin/xinliangnote/ginDemo2/router/v1"
	v2 "go-gin/xinliangnote/ginDemo2/router/v2"
	"go-gin/xinliangnote/ginDemo2/validator/member"
	"gopkg.in/go-playground/validator.v8"
)

func InitRouter(r *gin.Engine)  {
    //添加日志中间件
	r.Use(logger.LoggerToFile())



	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}
	// v2 版本
	GroupV2 := r.Group("/v2").Use(sign.Sign())
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid",member.NameValid)
	}



}
