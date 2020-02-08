package main

import (
	"fmt"
	"log"

	"github.com/kabuto412rock/microblog/model"

	"github.com/gin-gonic/gin"
)

type Env struct {
	mydb *model.MyDB
}

func main() {

	r := engine()
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

	r.GET("/", env.index)
	r.POST("login", env.login)
	return r
}

func (e *Env) login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userID, ok := e.mydb.GetUserID(username, password)
	if ok {
		// 成功登入
		c.JSON(200, gin.H{
			"Success": fmt.Sprintf("Hello, UserID:%d", userID),
		})
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
