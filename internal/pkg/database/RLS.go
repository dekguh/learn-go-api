package database

import (
	"log"

	"gorm.io/gorm"
)

func InitRLS(db *gorm.DB) {

	log.Println("RLS enabled")
}
