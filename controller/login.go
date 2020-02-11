package controller

import (
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (e *Env) Login(c *gin.Context) {
	userID := c.PostForm("userID")
	password := e.GetMD5Hash(c.PostForm("password"))

	session := sessions.Default(c)
	userID, ok := e.GetUserID(userID, password)
	if ok {
		// 成功登入
		session.Set(UserKey, userID)
		expire := 3600 * 8
		session.Options(sessions.Options{HttpOnly: true, MaxAge: expire})
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Fail": "Your session can't save",
			})
			return
		}
		cookie := http.Cookie{
			Name:    UserNameKey,
			Value:   username,
			Expires: time.Now().AddDate(0, 2, 1),
		}
		http.SetCookie(c.Writer, &cookie)
		c.Redirect(http.StatusFound, "list")
		// http.Redirect(c.Writer, c.Request, "list", http.StatusFound)
		c.Abort()
		return
	}
	// 登入失敗
	c.JSON(200, gin.H{
		"Login": "Fail",
	})
}
