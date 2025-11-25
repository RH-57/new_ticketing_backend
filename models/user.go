package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Password  string         `json:"password"`
	Role      string         `json:"role" gorm:"type:enum('superadmin','admin');default:'admin'"`
	Status    string         `json:"status" gorm:"type:enum('active', 'inactive');default:'active'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
