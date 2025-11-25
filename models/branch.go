package models

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	Id        uint           `json:"id" gorm:"primary_key"`
	Code      string         `json:"code" gorm:"unique;not null"`
	Name      string         `json:"name" gorm:"not null"`
	Divisions []Division     `json:"divisions,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
