package model

import "time"


type User struct{
	ID uint64 `gorm:"primayKey"`
	name string `gorm:"size:255;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	CreateAt time.Time
	UpdateAt time.Time 
}

