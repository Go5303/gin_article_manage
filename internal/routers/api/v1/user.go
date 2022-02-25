package v1

import (
	"blog-serice/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

type Data struct {
	CookieInfo interface{} `json:"cookieInfo"`
}

type Info struct {
	W http.ResponseWriter
	R *http.Request
}


var store = sessions.NewCookieStore([]byte("something-very-secret"))

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

/**
 设置session
 */
func SaveSession(c *gin.Context) {
	response := app.NewResponse(c)
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	//session 中做存储
	session.Values["foo"] = "bar"
	sessions.Save(c.Request, c.Writer)
	response.ToResponse(gin.H{"code":"0", "msg":"success"})
}
