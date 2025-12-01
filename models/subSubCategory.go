package models

import "time"

type SubSubCategory struct {
	Id            uint        `gorm:"primaryKey"`
	Name          string      `gorm:"size:255;not null"`
	Slug          string      `gorm:"size:255;unique"`
	SubCategoryId uint        `gorm:"not null"`
	SubCategory   SubCategory `gorm:"foreignKey:SubCategoryId"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
