package v1

import (
	"blog-serice/global"
	"blog-serice/internal/service"
	"blog-serice/pkg/app"
	"blog-serice/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {}

func NewArticle() Article {
	return Article{}
}

/**
 新增文章
 */
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	global.Logger.Infof("输入的参数为: %v", param)

	if !valid {
		global.Logger.Infof("参数校验异常，异常原因: %v", errs)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Infof("文章新增失败，失败原因: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	//文章新增成功
	response.ToResponse(gin.H{"code":"0", "msg":"success"})
}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
func (a Article) List(c *gin.Context) {}