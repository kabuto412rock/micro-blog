package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (e *Env) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(UserKey)
	session.Save()
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
