package models

import (
	"time"
)

type Category struct {
	Id          uint          `json:"id" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"not null"`
	Slug        string        `json:"slug" gorm:"unique;not null"`
	SubCategory []SubCategory `json:"sub_categories,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
