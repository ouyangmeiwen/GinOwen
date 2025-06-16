package manager

import (
	"GINOWEN/global"
	"GINOWEN/models"
	"GINOWEN/models/dto"
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"fmt"
	"strings"
)

type FlyReadAppManager struct {
}

func (f FlyReadAppManager) Hello(req request.HelloInput) (resp response.HelloResp) {

	var builder strings.Builder
	builder.WriteString(req.Name)
	fmt.Println(builder.String()) // 输出：数字是：123

	resp.Message = req.Name + ",Hello, FlyRead!"
	resp.Status = 200
	resp.Success = true
	return resp
}

func SetFlyReadSetting(req dto.FlyReadSetting, tenantid int) (err error) {
	// 这里可以添加逻辑来处理设置
	var sysconfig models.Systenantconfig
	err = global.OWEN_DB.Model(&models.Systenantconfig{}).Where("TenantId = ?", tenantid).First(&sysconfig).Error
	if err != nil {
		return err
	}
	if sysconfig.System == nil {

	}

	return nil
}
