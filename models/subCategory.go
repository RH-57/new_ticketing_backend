package models

import (
	"time"
)

type SubCategory struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	Slug       string    `json:"slug" gorm:"unique;not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   Category  `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
