package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/kabuto412rock/microblog/model"

	"github.com/gin-gonic/gin"
)

type Env struct {
	mydb *model.MyDB
}

const (
	USER_KEY = "user"
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

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 起始頁面(含登入頁面)
	r.GET("/", env.index)
	// 登入請求
	r.POST("login", env.login)

	// 列表介面
	authGroup := r.Group("/auth")
	authGroup.Use(AuthRequired)
	authGroup.GET("/list", env.articleList)
	return r
}

func (e *Env) login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	session := sessions.Default(c)
	userID, ok := e.mydb.GetUserID(username, password)
	if ok {
		// 成功登入
		session.Set("user", userID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Fail": "Your session can't save",
			})
			return
		}
		c.Redirect(http.StatusPermanentRedirect, "/auth/list")
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
	c.JSON(200, gin.H{"恭喜": fmt.Sprintf("成功登入:%d", userID)})
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
