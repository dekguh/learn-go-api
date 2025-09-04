package model

import "time"

type Todo struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	Title       string    `gorm:"not null;size:128"`
	Description string    `gorm:"not null;size:255"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
