package model

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	NameBook  string    `json:"name_book" gorm:"type:varchar(100);not null" binding:"required"`
	Author    string    `json:"author" gorm:"type:varchar(100);not null" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
