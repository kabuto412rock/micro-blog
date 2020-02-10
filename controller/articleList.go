package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (e *Env) ArticleList(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(UserKey)
	username, err := c.Cookie(UserNameKey)
	// 使用者根本沒登入 或登入時USER_NAME失效一樣錯誤返回
	if err != nil {
		NotFoundHandler(c)
		return
	}
	pageIndex, err := strconv.Atoi(c.Query("pageIndex"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	// 請求路徑的pageindex和pagesize有誤，幫忙重設定來到第一頁
	if err != nil || pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize > 25 || pageSize < 1 {
		pageSize = 5
	}
	page, ok := e.GetArticlePageByIndex(pageIndex, pageSize)
	if !ok {
		c.JSON(http.StatusExpectationFailed,
			gin.H{
				"error": fmt.Sprintf("User %d, you can't get all articles.", userID),
			})
		return
	}
	c.HTML(200, "articleList.html", gin.H{
		"username": username,
		"userID":   userID,
		"page":     page,
	})
}
