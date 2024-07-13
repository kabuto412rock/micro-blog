package model

type User struct {
	UserID   string    `gorm:"type:varchar(255);uniqueIndex"`
	Name     string    `gorm:"type:varchat(255)"`
	Password string    `gorm:"type:varchar(255)"`
	Token    string    `gorm:"type:varchar(255)"`
	Articles []Article `gorm:"foreignKey:UserID;references:UserID"`
	// `userID` varchar(30) NOT NULL DEFAULT '',
	// `name` varchar(255) DEFAULT NULL,
	// `password` varchar(255) DEFAULT NULL,
	// `token` varchar(255) DEFAULT NULL,
}

func (db MyDB) GetUserName(userID, password string) (username string, ok bool) {
	var user User
	result := db.Model(user).Where("user_id = ? AND password = ?", userID, password).First(&user)
	if result.Error != nil {
		return "", false
	}
	return user.Name, true
}

func (db MyDB) isUserIDValid(userID string) (ok bool) {
	// 使用者ID
	if len(userID) < 1 {
		return false
	}
	var count int64
	result := db.Model(&User{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil || count != 0 {
		return false
	}
	return true

}
func (db MyDB) CreateUser(userID, username, encodePassword string) (ok bool) {
	if len(username) < 1 {
		return false
	}
	if ok := db.isUserIDValid(userID); !ok {
		return false
	}
	user := User{
		UserID:   userID,
		Name:     username,
		Password: encodePassword,
	}
	result := db.Create(&user)
	if result.Error != nil || result.RowsAffected < 1 {
		return false
	}
	return true
}
