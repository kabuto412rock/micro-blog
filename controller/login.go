package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (e *Env) Login(c *gin.Context) {
	userID := c.PostForm("userID")
	password := e.GetMD5Hash(c.PostForm("password"))

	session := sessions.Default(c)
	username, ok := e.GetUserName(userID, password)
	// 成功登入
	if ok {
		session.Set(UserKey, userID)
		expire := 3600 * 8
		session.Options(sessions.Options{HttpOnly: true, MaxAge: expire})
		if err := session.Save(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Fail": "Your session can't save",
			})
			return
		}
		// cookie := http.Cookie{
		// 	Name:    UserNameKey,
		// 	Value:   username,
		// 	Expires: time.Now().AddDate(0, 2, 1),
		// }

		// http.SetCookie(c.Writer, &cookie)
		c.SetCookie(UserNameKey, username, 36000, "/", Domain, false, true)
		c.Redirect(http.StatusFound, "list")

		c.Abort()
		return
	}
	// 登入失敗
	c.JSON(200, gin.H{
		"Login": "Fail",
	})
}
