package model

import (
	"database/sql"
	"fmt"

	// Mysql的Driver
	_ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	*sql.DB
}

/*New 產生一個DB實例*/
func New(dbUser, dbPassword, dbLocalhost, dbPort, dbName string) (*MyDB, error) {

	// 連接本地的MySQL資料庫
	db, err := connectMysql(dbUser,
		dbPassword,
		dbLocalhost,
		dbPort,
		dbName)

	return &MyDB{db}, err
}

// Close 關閉MyDB內部的*sql.db實例
func (mydb MyDB) Close() error {
	return mydb.Close()
}

// 連接到本地端名為dbName的資料庫
func connectMysql(userName, userPassword, dbLocalhost, port, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		userName,
		userPassword,
		dbLocalhost,
		port,
		dbName)

	return sql.Open("mysql", dataSourceName)
}
