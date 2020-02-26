# MicroBlog

## a bolg server using Gin framework written in Go and MySQL database

## 功能實現
1. 文章CRUD(Create, Read, Update, Delete)
2. 基本使用者登入、登出、session驗證
3. 把密碼加密存到資料庫
4. 註冊帳號(無Email驗證，單純新增帳號)

## 未來目標

1. 將PostForm 改成使用AJax的方式更新資料。
2. 使用JWT token授權取得資料
3. 註冊新帳號(預計使用Email驗證)
4. 將model介面都改成REST API形式配合第一項執行

## 如何啟動
1. 將專案複製到電腦中
 ```bash
 $ git clone https://github.com/kabuto412rock/micro-blog.git
 ```
2. 啟動你自己的MySQL Server
```bash
# 因人而異，我是用homebrew安裝mysql server
$ mysql.server start
``` 
3. 欲修改Server和DB的設定，請查看config.yaml
```yaml
# Server configuations
server:
  host: "localhost" 
  port: 8080

# MySQL Database configuations
database:
  user: "dbuser" # DB User name
  pass: "Ej3yj/ru8@" # DB User password
  dbname: "UserDB" # DB name
  host: "127.0.0.1" # DB Server's host 
  port: 3306 # DB Server's port
```
4. 新建MySQL資料表 
##### Article資料表
```sql
CREATE TABLE `Article` (
  `articleID` int NOT NULL AUTO_INCREMENT,
  `userID` varchar(30) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `editTime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`articleID`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
##### User資料表
```sql
CREATE TABLE `User` (
  `userID` varchar(30) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
5. 啟動Golang Server
```bash
$ go run main.go
```

## 介面說明