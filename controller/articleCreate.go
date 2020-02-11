package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kabuto412rock/microblog/model"
)

func (e Env) ArticleCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	session := sessions.Default(c)
	userID, ok := session.Get(UserKey).(string)

	if title == "" || content == "" || !ok {
		c.Redirect(302, "list")
		c.Abort()
		return
	}
	a := model.Article{Title: title, Content: content, UserID: userID}
	if ok := e.InsertArticle(a); !ok {
		c.JSON(404, gin.H{
			"錯囉": "InsertArticle",
		})
		return
	}
	c.Redirect(http.StatusFound, "list")
	c.Abort()
}
