package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kabuto412rock/microblog/model"
)

func (e Env) ArticleEdit(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	session := sessions.Default(c)
	userID, ok := session.Get(UserKey).(string)
	articleID, err := strconv.Atoi(c.Query("articleID"))

	if title == "" || content == "" || !ok || err != nil {
		c.Redirect(302, "list")
		c.Abort()
		return
	}
	a := model.Article{Title: title, Content: content, ArticleID: articleID, UserID: userID}

	ok = e.UpdateArticle(a) // ommit the ok
	if !ok {
		c.JSON(404, gin.H{
			"error": "為什麼是卡在這？！",
			"article": gin.H{
				"Title":     title,
				"articleID": articleID,
				"content":   content,
				"UserID":    userID,
			},
		})
		c.Abort()
		return
	}
	c.Redirect(http.StatusFound, "list")
	c.Abort()
}
