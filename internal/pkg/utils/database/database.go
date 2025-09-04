package dbutils

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SetCurrentUserId(db *gorm.DB, userID uint) {
	sql := fmt.Sprintf("SET app.current_user_id = %d", userID)
	log.Println("sql: ", sql)
	if err := db.Exec(sql).Error; err != nil {
		log.Println("failed to set current user id: ", err)
	}
	log.Println("app.current_user_id: ", userID)
}
