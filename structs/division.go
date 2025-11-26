package structs

type DivisionResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	BranchId   uint   `json:"branch_id"`
	BranchCode string `json:"branch_code"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type DivisionCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	BranchId uint   `json:"branch_id" binding:"required"`
}

type DivisionUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	BranchId uint   `json:"branch_id"`
}

type DivisionWithBranchResponse struct {
	Id     uint            `json:"id"`
	Name   string          `json:"name"`
	Branch BranchSimpleRes `json:"branch"`
}

type BranchSimpleRes struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
