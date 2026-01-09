package model

import "time"


type User struct{
	ID uint64 `gorm:"primayKey"`
	Email string `gorm:"size:255;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time 
}

