package v1

import (
	"blog-serice/global"
	"blog-serice/internal/service"
	"blog-serice/pkg/app"
	"blog-serice/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {}

func NewTag() Tag {
	return Tag{}
}

/**
 创建标签
 */
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Infof("输入的参数为: %v", param)
	if !valid {
		global.Logger.Infof("参数校验异常，异常原因: %v", errs)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Infof("标签新增失败，失败原因: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//标签新增成功
	response.ToResponse(gin.H{"code":"0","msg":"success"})
}

/**
 获取标签列表
 */
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid { //参数校验不通过
		global.Logger.Infof("参数校验异常，异常原因: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Infof("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Infof("svc.ListTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

/**
 删除标签
 */
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Infof("输入的参数为: %v", param)
	if !valid {
		global.Logger.Infof("参数校验异常，异常原因: %v", errs)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Infof("标签删除失败，失败原因: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{"code":"0","msg":"success"})  //标签删除成功
}

/**
 更新标签
 */
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Info("输入的参数为: %v", param)
	if !valid {
		global.Logger.Info("参数校验异常，异常原因: %v", errs)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Info("标签修改失败，失败原因: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{"code":0,"msg":"success"})  //标签更新成功
}

