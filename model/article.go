package model

import "time"

/*Article 資料表是MySQL資料庫UserDB中的資料表，
儲存使用者撰寫的文章*/
type Article struct {
	ArticleID int
	UserID    int
	Title     string
	Content   string
	EditTime  time.Time
	username  string
}

func (db MyDB) GetArticlesCount() (count int, ok bool) {
	row := db.QueryRow("SELECT COUNT(*) FROM Article")
	if err := row.Scan(&count); err != nil {
		return count, false
	}
	return count, true
}

// GetArticlesByPage 利用取得 某一頁的文章
// page 指定要的頁數文章，從1開始
// onePageSize 一個頁面所呈現的文章個數，Ex: onePageSize=10，回傳的articles數量最多不超過10個
func (db MyDB) GetArticlesByPage(page int, onePageSize int) (articles []Article, ok bool) {
	indexStart := (page - 1) * onePageSize
	rows, err := db.Query("SELECT articleID, userID, title, content, editTime FROM Article LIMIT ? OFFSET ?", onePageSize, indexStart)
	if err != nil {
		return articles, false
	}
	var a Article
	for rows.Next() {
		err := rows.Scan(&a.ArticleID, &a.UserID, &a.Title, &a.Content, &a.EditTime)
		if err != nil {
			return articles, false
		}
		articles = append(articles, a)
	}
	return articles, true
}

// 取得所有的Article
func (db MyDB) GetAllArticles() (articles []Article, ok bool) {
	rows, err := db.Query("SELECT articleID, userID, title, content, editTime FROM Article")
	if err != nil || rows == nil {
		return nil, false
	}

	var one Article
	for rows.Next() {
		rows.Scan(&one.ArticleID, &one.UserID, &one.Title, &one.Content, &one.EditTime)
		articles = append(articles, one)
	}
	return articles, true
}

/*插入一個Article*/
func (db MyDB) insertArticle(a Article) (ok bool) {
	result, err := db.Exec(`
	INSERT INTO Article(userID,title, content)
	Values(?, ?, ?)
	`)
	if err != nil {
		return false
	}
	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return false
	}
	return true
}

// 更新一個Article
func (db MyDB) updateArticle(a Article) (ok bool) {
	result, err := db.Exec(`
	UPDATE Article
	SET title=?, content=?
	WHERE articleID=?
	`, a.Title, a.Content, a.ArticleID)

	if err != nil {
		return false
	}
	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return false
	}
	return true
}

// 刪除一個Article
func (db MyDB) deleteArticle(articleID int) (ok bool) {
	result, err := db.Exec(
		`DELETE FROM Article WHERE articleID=?`, articleID)
	if err != nil {
		return false
	}
	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return false
	}
	return true
}
