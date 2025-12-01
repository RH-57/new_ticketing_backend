package structs

import "time"

type SubCategoryResponse struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	CategoryId   uint      `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SubCategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type SubCategoryUpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	CategoryId uint   `json:"category_id"`
}
