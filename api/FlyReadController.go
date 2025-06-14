package api

import (
	"GINOWEN/global"
	"GINOWEN/models/request"
	"GINOWEN/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FlyReadController struct {
}

// Hello
// @Tags     FlyRead
// @Summary  查询书架
// @Produce  application/json
// @Param    data  query     request.HelloInput 			true  "参数"
// @Success  200   {object}  utils.Response{data=[]response.HelloResp,msg=string}  "返回清单"
// @Security BearerAuth
// @Router   /api/services/app/FlyRead/Hello [get]
func (b *FlyReadController) Hello(c *gin.Context) {
	// 打印查询字符串
	queryStr := c.Request.URL.RawQuery
	fmt.Println("Query String:", queryStr)
	//post
	var req request.HelloInput
	//err := c.ShouldBindJSON(&req)
	//get
	err := c.ShouldBindQuery(&req) //大小写敏感

	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	list, err := ServicesGroup.flyreadAppService.Hello(req)
	if err != nil {
		global.OWEN_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		utils.FailWithMessage("获取失败!"+err.Error(), c)
		return
	}
	utils.OkWithDetailed(list, "获取成功", c)
}
