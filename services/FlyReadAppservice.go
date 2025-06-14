package services

import (
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"GINOWEN/services/manager"
)

type FlyReadAppService struct {
}

var flyReadAppManager = &manager.FlyReadAppManager{}

func (f *FlyReadAppService) Hello(req request.HelloInput) (resp response.HelloResp, err error) {
	return flyReadAppManager.Hello(req), nil
}
