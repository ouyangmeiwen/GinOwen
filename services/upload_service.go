package services

import (
	"GINOWEN/models/request"
	"GINOWEN/models/response"
	"fmt"
)

type UploadfileService struct {
}

func (*UploadfileService) UploadFileHandler(req request.UploadRequest) (resp response.UploadResponse, err error) {
	fmt.Printf("接收到上传信息 %s, %s", req.Name, req.Description)
	return resp, err
}
