package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (e *Env) Index(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(UserKey)
	
	// 前往登入頁面
	if userID == nil || userID == "" {
		c.HTML(200, "login.html", nil)
		return
	}
	// 已登入，前往文章列表頁面
	c.Redirect(http.StatusFound, "list")
}
