package model

func (db MyDB) GetUserID(username, password string) (userID int, ok bool) {

	row := db.QueryRow("SELECT userID from User WHERE name = ? and  password=?", username, password)
	if err := row.Scan(&userID); err != nil {
		return userID, false
	}
	return userID, true
}
func (db MyDB) isUserNameValid(username string) (ok bool) {
	row := db.QueryRow(`
	SELECT COUNT(*) FROM User WHERE name=?`, username)
	var count int
	err := row.Scan(&count)
	if err != nil || count != 0 {
		return false
	}
	return true

}
func (db MyDB) CreateUser(username, encodePassword string) (ok bool) {
	if ok := db.isUserNameValid(username); !ok {
		return false
	}
	result, err := db.Exec(`
	INSERT INTO User(name, password)
	Values(?, ?)
	`, username, encodePassword)
	if err != nil {
		return false
	}
	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return false
	}
	return true
}
