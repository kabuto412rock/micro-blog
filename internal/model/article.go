package model

import (
	"time"
)

type ArticlePage struct {
	Articles           []ArticleResult
	CurrentIndex       int
	AnotherPageIndices []int
	PageSize           int
}

/*
Article 資料表是MySQL資料庫UserDB中的資料表，
儲存使用者撰寫的文章
*/
type Article struct {
	ArticleID int    `gorm:"primaryKey"`
	UserID    string `gorm:"type:varchar(255);index"`
	Title     string
	Content   string
	EditTime  time.Time
}
type ArticleResult struct {
	ArticleID int       `gorm:"article_id"`
	UserID    string    `gorm:"user_id"`
	Title     string    `gorm:"title"`
	Content   string    `gorm:"content"`
	EditTime  time.Time `gorm:"edit_time"`
	UserName  string    `gorm:"user_name"`
}

func (db MyDB) GetArticlesCount() (count int64, ok bool) {
	result := db.Model(&Article{}).Count(&count)
	if result.Error != nil {
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
	var articles []ArticleResult

	// 執行查詢
	err := db.Model(&Article{}).
		Select("articles.article_id, articles.user_id, articles.title, articles.content, articles.edit_time, users.name as user_name").
		Joins("left join users on articles.user_id = users.user_id").
		Order("articles.edit_time desc").
		Limit(onePageSize).
		Offset(indexStart).
		Scan(&articles).Error
	if err != nil {
		return nil, false
	}
	// 取得資料庫中Article的數量並計算總共會產生幾個頁面
	allArticleCount, _ := db.GetArticlesCount()
	allPageCount := int(allArticleCount) / onePageSize
	if int(allArticleCount)%onePageSize != 0 {
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
	article := &Article{
		UserID:   a.UserID,
		Title:    a.Title,
		Content:  a.Content,
		EditTime: time.Now(),
	}
	result := db.Create(article)
	if result.Error != nil || result.RowsAffected < 1 {
		return false
	}
	return true
}

// 更新一個Article
func (db MyDB) UpdateArticle(a Article) (ok bool) {
	article := &Article{
		Title:    a.Title,
		Content:  a.Content,
		EditTime: time.Now(),
	}
	result := db.Model(&Article{}).
		Where("article_id = ? AND user_id = ?", a.ArticleID, a.UserID).
		Updates(article)
	if result.Error != nil || result.RowsAffected < 1 {
		return false
	}
	return true
}

// 刪除一個Article
func (db MyDB) DeleteArticle(articleID int, userID string) (ok bool) {
	result := db.Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&Article{})
	// result, err := db.Exec(
	// 	`DELETE FROM Article WHERE articleID=? AND userID = ?`, articleID, userID)
	// if err != nil {
	// 	return false
	// }
	// if rows, err := result.RowsAffected(); err != nil || rows < 1 {
	// 	return false
	// }
	if result.Error != nil || result.RowsAffected < 1 {
		return false
	}
	return true
}
