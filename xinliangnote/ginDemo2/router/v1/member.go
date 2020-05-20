package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin/xinliangnote/ginDemo2/entity"
	"net/http"
)

func AddMember(c *gin.Context) {

	res := entity.Result{}
	mem := entity.Member{}


	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_ERROR)
		res.SetMessage("参数验证错误")
		//res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}
	//@todo 处理业务(下次再分享)
	//@todo ...

    //返回
	data := map[string]interface{}{
		"name" : mem.Name,
		"age"  : mem.Age,
	}
	res.SetCode(entity.CODE_SUCCESS)
	res.SetData(data)
	c.JSON(http.StatusOK, res)
}
