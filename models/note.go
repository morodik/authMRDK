package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
