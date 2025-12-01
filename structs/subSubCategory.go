package structs

type SubSubCategoryResponse struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	SubCategoryId   uint   `json:"sub_category_id"`
	SubCategoryName string `json:"sub_category_name"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type SubSubCategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type SubSubCategoryUpdateRequest struct {
	Name          string `json:"name" binding:"required"`
	SubCategoryId uint   `json:"sub_category_id"` // optional, boleh pindahkan ke subcategory lain
}
