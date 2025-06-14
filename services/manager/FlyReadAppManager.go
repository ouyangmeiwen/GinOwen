package manager

import (
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
