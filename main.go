package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/kabuto412rock/microblog/model"

	"github.com/gin-gonic/gin"
)

type Env struct {
	*model.MyDB
}

const (
	USER_KEY      = "user"
	USER_NAME_KEY = "user_name"
)

func main() {
	// 建立一個新Http服務 r
	r := engine()

	// r 開始在本地服務Port: 8080
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal("r.Run's error:", err)
	}
}

func engine() *gin.Engine {
	mydb, err := model.New()
	if err != nil {
		log.Fatal("mydb's error:", err)
	}
	env := &Env{mydb}

	r := gin.New()
	r.LoadHTMLGlob("template/*")

	// 使用紀錄CookieSession的中介層
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	/* 無需驗證的路由 */
	// 起始頁面(即登入頁面)
	r.GET("/", env.index)
	// 登入請求
	r.POST("login", env.login)

	// 登入驗證的中介層(以是否存在session辨識使用者是否已登入)
	r.Use(AuthRequired)

	// 文章列表頁面
	r.GET("list", env.articleList)
	return r
}

func (e *Env) login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	session := sessions.Default(c)
	userID, ok := e.GetUserID(username, password)
	if ok {
		// 成功登入
		session.Set("user", userID)
		expire := 3600 * 8
		session.Options(sessions.Options{HttpOnly: true, MaxAge: expire})
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Fail": "Your session can't save",
			})
			return
		}
		cookie := http.Cookie{
			Name:    USER_NAME_KEY,
			Value:   username,
			Expires: time.Now().AddDate(0, 0, 1),
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
func (e *Env) index(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (e *Env) articleList(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(USER_KEY)
	username, err := c.Cookie(USER_NAME_KEY)
	if err != nil {
		username = "Undefined User"
	}
	articles, err := e.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusExpectationFailed,
			gin.H{
				"error": fmt.Sprintf("User %d, you can't get all articles because %s", userID, err),
			})
		return
	}
	c.HTML(200, "articleList.html", gin.H{
		"username": username,
		"articles": articles,
	})
}

// AuthRequired 驗證使用者登入的中介層(採用Cookie Session)
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(USER_KEY)

	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})

		return
	}
	// 有登入過的Session紀錄，繼續執行
	c.Next()
}
