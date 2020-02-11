package model

func (db MyDB) GetUserName(userID, password string) (username int, ok bool) {

	row := db.QueryRow("SELECT userID from User WHERE userID = ? and  password=?", userID, password)
	if err := row.Scan(&username); err != nil {
		return username, false
	}
	return username, true
}

func (db MyDB) isUserIDValid(userID string) (ok bool) {
	// 使用者ID 至少十碼
	if len(userID) < 10 {
		return false
	}
	row := db.QueryRow(`
	SELECT COUNT(*) FROM User WHERE userID=?`, userID)
	var count int
	err := row.Scan(&count)
	if err != nil || count != 0 {
		return false
	}
	return true

}
func (db MyDB) CreateUser(userID, username, encodePassword string) (ok bool) {
	if ok := db.isUserIDValid(userID); !ok {
		return false
	}
	if len(username) < 3 {
		return false
	}
	result, err := db.Exec(`
	INSERT INTO User(userID, name, password)
	Values(?, ?, ?)
	`, userID, username, encodePassword)
	if err != nil {
		return false
	}
	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return false
	}
	return true
}
