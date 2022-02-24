package routers

import (
	"blog-serice/internal/middleware"
	v1 "blog-serice/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())  //注入日志
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	tag := v1.NewTag()

	//路由管理
	v1Api := r.Group("/api/v1")
	{
		//标签管理
		v1Api.GET("/tags/list", tag.List)
		v1Api.POST("/tags", tag.Create)
		v1Api.POST("/tags/delete", tag.Delete)
		v1Api.POST("/tags/update", tag.Update)
	}
	return r
}
