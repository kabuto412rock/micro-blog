package controller

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kabuto412rock/microblog/model"
)

type Env struct {
	*model.MyDB
	salt []byte
}

const (
	UserKey     = "user_id"   // Key for Session & Cookie
	UserNameKey = "user_name" // Key for Cookie
	Domain      = "localhost" // WEB Server's domain
	// 底下的常數自行調整。
	DBUser      = "dbuser"     // MySQL's User name
	DBPassword  = "Ej3yj/ru8@" // MySQL's User password
	DBName      = "UserDB"     // MySQL's DB Name
	DBLocalhost = "127.0.0.1"  // MySQL Server's IP Address
	DBport      = "3306"       // MySQL port
)

func NewEnv(salt []byte) *Env {
	mydb, err := model.New(
		DBUser,
		DBPassword,
		DBLocalhost,
		DBport,
		DBName,
	)
	if err != nil {
		log.Fatal("mydb's error:", err)
	}

	env := &Env{MyDB: mydb, salt: salt}
	return env
}

// AuthRequired 驗證使用者登入的中介層(採用Cookie Session)
func (e Env) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserKey)

	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// 有登入過的Session紀錄，繼續執行
	c.Next()
}

func (e Env) GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(e.salt))
}

func NotFoundHandler(c *gin.Context, err error) {
	c.JSON(404, gin.H{
		"error": "兄弟，你的錯誤是:" + err.Error(),
	})
	c.Abort()
}
