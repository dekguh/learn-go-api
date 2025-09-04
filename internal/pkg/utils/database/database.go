package dbutils

import "gorm.io/gorm"

func SetCurrentUserId(db *gorm.DB, userID uint) {
	db.Exec("SET app.current_user = ?", userID)
}
