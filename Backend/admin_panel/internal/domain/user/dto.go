package user

type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email,max=100"`
	Role      string `json:"role" validate:"required,oneof=Admin Editor Viewer"`
	IsActive  *bool  `json:"isActive" validate:"required"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"lastName" validate:"omitempty,min=2,max=50"`
	Email     *string `json:"email" validate:"omitempty,email,max=100"`
	Role      *string `json:"role" validate:"omitempty,oneof=Admin Editor Viewer"`
	IsActive  *bool   `json:"isActive" validate:"omitempty"`
}

type Pagination struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}
