package dto

type UserRequest struct {
	User     string `validate:"required" json:"user"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Role     string `json:"role"`
}
