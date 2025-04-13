package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID  uint   `gorm:"not null" json:"user_id"`
	Title   string `gorm:"not null" json:"title"`
	Content string `json:"content"`
}
