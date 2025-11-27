package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	Id         uint           `json:"id" gorm:"primaryKey"`
	Name       string         `json:"name" gorm:"not null"`
	DivisionId uint           `json:"division_id" gorm:"not null"`
	Division   Division       `json:"-"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
