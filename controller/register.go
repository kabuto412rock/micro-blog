package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterGET 返回一個註冊頁面
func (e Env) RegisterGET(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

// RegisterPOST 處理由reigster.html發送的POST註冊請求
func (e Env) RegisterPOST(c *gin.Context) {
	// 取得 Post請求內的帳號
	userID := c.PostForm("userID")
	// 取得 Post請求內的暱稱
	username := c.PostForm("username")
	// 取得 被MD5加密後的Post請求內的密碼
	password := e.GetMD5Hash(c.PostForm("password"))

	// 根據帳號、密碼新增一個User
	if ok := e.CreateUser(userID, username, password); !ok {
		c.AbortWithStatusJSON(http.StatusPermanentRedirect,
			gin.H{
				"error": "無法新建此帳號",
			})
		return
	}
	// 成功建立User帳號，轉址到主頁面讓使用者自己登入。
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
