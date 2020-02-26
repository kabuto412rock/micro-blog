package model

import (
	"database/sql"
	"fmt"

	cfg "github.com/kabuto412rock/microblog/config"

	// Mysql的Driver
	_ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	*sql.DB
}

/*New 產生一個MyDB實例*/
func New(config *cfg.Config) (*MyDB, error) {
	dbConfig := &config.Database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName)

	// 連接本地的MySQL資料庫
	db, err := sql.Open("mysql", dataSourceName)
	// 回傳一個
	return &MyDB{db}, err
}

// Close 關閉MyDB內部的*sql.db實例
func (mydb MyDB) Close() error {
	return mydb.Close()
}
