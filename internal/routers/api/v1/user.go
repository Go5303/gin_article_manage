package v1

import (
	"blog-serice/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	CookieInfo interface{} `json:"cookieInfo"`
}

type Info struct {
	W http.ResponseWriter
	R *http.Request
}

/**
 设置cookie
 */
func SetCookie(c *gin.Context) {
	response := app.NewResponse(c)
	_, err := c.Cookie("username")
	if err != nil {
		c.SetCookie("username", "zhanghai", 3600, "/", "*",false, true)
	}
	response.ToResponse(gin.H{"code":"0", "msg":"success", "data":Data{CookieInfo: "张海"}})
}
