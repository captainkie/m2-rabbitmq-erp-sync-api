package request

type CreateUsersRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Email    string `validate:"required,email,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	// Role     string `json:"role"`
	// Status   string `json:"status"`
	// Created  string `json:"created"`
	// Updated  string `json:"updated"`
}

type UpdateUsersRequest struct {
	Password string `validate:"required,min=1,max=100" json:"password"`
	Role     string `validate:"required,min=1,max=100" json:"role"`
	Status   string `validate:"required,min=1,max=100" json:"status"`
}

type LoginRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
