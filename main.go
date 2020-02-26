package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/kabuto412rock/microblog/config"
	"github.com/kabuto412rock/microblog/controller"
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
	// 讀取config.yaml的設定
	myConfig := config.ReadConfig()
	
	env := controller.NewEnv(myConfig)

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

	r.Static("/js/", "static/js/")
	r.Static("/css/", "static/css/")
	r.Static("/img/", "static/img/")
	// 註冊請求 Get->回報頁面, Post->
	r.GET("register", env.RegisterGET)
	r.POST("register", env.RegisterPOST)

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
