package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Name      string `gorm:"size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
