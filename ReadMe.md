# MicroBlog

## a bolg server using Gin framework written in Go and MySQL database

## 完成

基本的CRUD完成，但並沒有Restful API。  
基本使用者登入、登出、session驗證
把密碼加密存到資料庫

## 待完成

2. 註冊新帳號(預計使用Email驗證)
3. 文章排版美化(至少操作畫面要平順，各個區塊明確切割)
4. 將model介面都改成REST API形式配合第一項執行。

## 未來目標

1. 將PostForm 改成使用AJax的方式更新資料。
2. 使用JWT token授權取得資料