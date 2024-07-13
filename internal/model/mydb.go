package model

import (
	"fmt"

	cfg "github.com/kabuto412rock/microblog/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Mysql的Driver
	_ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	*gorm.DB
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
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Article{}, &User{})
	return &MyDB{db}, err
}
