package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID      uint       `gorm:"not null" json:"user_id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	Status      bool       `gorm:"default:false" json:"status"`
}
