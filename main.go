package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := engine()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal("r.Run's error:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	r.GET("/", index)
	return r
}
func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"index": "這是主頁面",
	})
}
