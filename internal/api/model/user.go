package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;size:255;not null"`
	Password  string `json:"password,omitempty" gorm:"size:255;not null"`
	Name      string `gorm:"size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDetailNoPasswordResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
