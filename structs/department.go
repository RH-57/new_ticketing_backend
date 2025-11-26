package structs

type DepartmentResponse struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	DivisionId   uint   `json:"division_id"`
	DivisionName string `json:"division_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type DepartmentCreateRequest struct {
	Name       string `json:"name" binding:"required"`
	DivisionId uint   `json:"division_id" binding:"required"`
}

type DepartmentUpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	DivisionId uint   `json:"division_id"`
}
