package model

import "time"

type ArticlePage struct {
	Articles           []Article
	CurrentIndex       int
	AnotherPageIndices []int
	PageSize           int
}

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
// currentPageIndex 指定要的頁數文章，從1開始
// onePageSize 一個頁面所呈現的文章個數，Ex: onePageSize=10，回傳的articles數量最多不超過10個
func (db MyDB) GetArticlePageByIndex(currentPageIndex int, onePageSize int) (page *ArticlePage, ok bool) {
	if onePageSize < 1 || currentPageIndex < 1 {
		return nil, false
	}
	// ex: onePageSize = 5, currentPageIndex
	indexStart := (currentPageIndex - 1) * onePageSize
	rows, err := db.Query("SELECT articleID, userID, title, content, editTime FROM Article ORDER BY editTime DESC LIMIT ? OFFSET ?", onePageSize, indexStart)
	var articles []Article
	if err != nil {
		return nil, false
	}
	var articleCount int = 0
	var a Article
	// 依序取得Article放入articles並計數有幾篇(articleCount)文章
	for rows.Next() {
		err := rows.Scan(&a.ArticleID, &a.UserID, &a.Title, &a.Content, &a.EditTime)
		if err != nil {
			return nil, false
		}
		articles = append(articles, a)
		articleCount++
	}
	// 取得資料庫中Article的數量並計算總共會產生幾個頁面
	allArticleCount, ok := db.GetArticlesCount()
	allPageCount := allArticleCount / onePageSize
	if allArticleCount%onePageSize != 0 {
		allPageCount++
	}
	// 產生ArticlePage底部文章列表的連結索引值
	var anotherPageIndices []int

	// 底部文章連結最小頁數的索引
	var anotherPageIndexStart int = currentPageIndex - 2
	if anotherPageIndexStart < 1 {
		anotherPageIndexStart = 1
	}
	// 底部文章連結最大頁數的索引
	var anotherPageIndexEnd int = currentPageIndex + 2
	if anotherPageIndexEnd > allPageCount {
		anotherPageIndexEnd = allPageCount
	}
	for i := anotherPageIndexStart; i <= anotherPageIndexEnd; i++ {
		anotherPageIndices = append(anotherPageIndices, i)
	}

	// 回傳的文章頁面資料
	page = &ArticlePage{
		Articles:           articles,
		CurrentIndex:       currentPageIndex,
		AnotherPageIndices: anotherPageIndices,
		PageSize:           onePageSize,
	}
	return page, true
}

/*插入(新增)一個Article*/
func (db MyDB) InsertArticle(a Article) (ok bool) {
	result, err := db.Exec(`
	INSERT INTO Article(userID, title, content, editTime)
	Values(?, ?, ?, Now())
	`, a.UserID, a.Title, a.Content)
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
