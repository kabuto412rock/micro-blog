package model

import (
	"database/sql"
	"fmt"
	"time"

	// Mysql的Driver
	_ "github.com/go-sql-driver/mysql"
)

/*Article 資料表是MySQL資料庫UserDB中的資料表，
儲存使用者撰寫的文章*/
type Article struct {
	ArticleID int
	UserID    int
	Title     string
	Content   string
	EditTime  time.Time
}
type MyDB struct {
	*sql.DB
}

/*New 產生一個DB實例*/
func New() (*MyDB, error) {
	user := "dbuser"         // 帳號
	password := "Ej3yj/ru8@" // 密碼
	dbName := "UserDB"       // db名稱

	// 連接本地的MySQL資料庫
	db, err := connectMysql(user, password, dbName)

	return &MyDB{db}, err
}

// Close 關閉MyDB內部的*sql.db實例
func (mydb MyDB) Close() error {
	return mydb.Close()
}

// 連接到本地端名為dbName的資料庫
func connectMysql(userName, userPassword, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", userName, userPassword, dbName)

	return sql.Open("mysql", dataSourceName)
}

