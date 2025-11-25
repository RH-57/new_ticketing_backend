package models

import (
	"time"

	"gorm.io/gorm"
)

type Division struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	BranchId  uint           `json:"branch_id" gorm:"not null"` // Foreign Key
	Branch    Branch         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
