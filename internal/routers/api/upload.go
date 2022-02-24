package api

import (
	"blog-serice/global"
	"blog-serice/internal/service"
	"blog-serice/pkg/app"
	"blog-serice/pkg/convert"
	"blog-serice/pkg/errcode"
	"blog-serice/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context)  {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Info(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{"file_access_url":fileInfo.AccessUrl})
}
