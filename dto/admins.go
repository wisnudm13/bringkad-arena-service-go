package dto

type RegisterAdminRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

type UpdateAdminRequest struct {
	Email           string `json:"email"`
	IsActive        bool   `json:"is_active"`
	Password        string `json:"password" binding:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginAdminRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetAdminResponse struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
