package model

type Task struct{
	ID uint64 `gorm:"primaryKey"`
	Title string `gorm:"size:255;not null"`
}

