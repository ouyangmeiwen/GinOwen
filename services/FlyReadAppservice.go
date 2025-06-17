package services

import (
	"GINOWEN/models/request"
	"GINOWEN/models/response"
)

type FlyReadAppService struct {
}

func (f *FlyReadAppService) Hello(req request.HelloInput) (resp response.HelloResp, err error) {
	return ManagerGroup.frymanager.Hello(req), nil
}
