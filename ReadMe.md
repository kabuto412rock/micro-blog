# MicroBlog

## a bolg server using Gin framework written in Go and MySQL database

## 功能實現

1. 文章CRUD(Create, Read, Update, Delete)
2. 基本使用者登入、登出、session驗證
3. 把密碼加密存到資料庫
4. 註冊帳號(無Email驗證，單純新增帳號)
5. 基本防範XSS、CSRF攻擊
## 未來目標
1. 將PostForm 改成使用AJax的方式更新資料。
2. 使用JWT token授權取得資料
3. 註冊新帳號(預計使用Email驗證)
4. 將model介面都改成REST API形式配合第一項執行

## 開發環境要求
- 安裝[Docker Desktop](https://www.docker.com/products/docker-desktop/)
- 開發編輯器[VSCode](https://code.visualstudio.com/)安裝插件[Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) (讀取.devcontainer設定)

## 啟動專案
0. 將專案複製到電腦中
    ```bash
    git clone https://github.com/kabuto412rock/micro-blog.git 
    ```
1. 使用VSCode開啟專案，按下快捷鍵`F1`，選擇`dev containers: Reopen in Container`
2. 上述內容成功後進到容器，執行以下指令
    ```bash
    cd micro-blog  # 進到專案資料夾
    go build ./cmd/micro-blog # 將程式碼編譯成執行檔
    ./micro-blog # 啟動Server:8080
    ```
> 欲修改Server和DB的設定，請查看[config.yaml](./config.yaml)  
> 已改為使用VSCode的 devcontainer容器啟動

## 介面說明