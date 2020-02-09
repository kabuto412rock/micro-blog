package model

func (db MyDB) GetUserID(username, password string) (userID int, ok bool) {
	
	row := db.QueryRow("SELECT userID from User WHERE name = ? and  password=?", username, password)
	if err := row.Scan(&userID); err != nil {
		return userID, false
	}
	return userID, true
}
