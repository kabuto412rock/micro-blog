package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

func (e Env) ArticleDelete(c *gin.Context) {
	session := sessions.Default(c)
	userID, ok := session.Get(UserKey).(int)

	articleID, err := strconv.Atoi(c.Query("articleID"))
	if !ok || err != nil {
		c.JSON(404, gin.H{"error": "ArticleDelete有錯"})
		c.Abort()
		return
	}
	e.DeleteArticle(articleID, userID)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
