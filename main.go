package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kabuto412rock/microblog/controller"
	"github.com/kabuto412rock/microblog/model"
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

	env := &controller.Env{mydb}

	r := gin.New()
	r.LoadHTMLGlob("template/*")

	// 使用紀錄CookieSession的中介層
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	/* 無需驗證的路由 */
	// 起始頁面(即登入頁面)
	r.GET("/", env.Index)
	// 登入請求
	r.POST("login", env.Login)

	// 登入驗證的中介層(以是否存在session辨識使用者是否已登入)
	r.Use(env.AuthRequired)

	// 登出
	r.GET("logout", env.Logout)

	// 文章列表頁面
	r.GET("list", env.ArticleList)

	// 新建文章
	r.POST("create", env.ArticleCreate)

	// 刪除文章
	r.POST("delete", env.ArticleDelete)

	// 更新文章
	r.POST("update", env.ArticleEdit)
	return r
}
