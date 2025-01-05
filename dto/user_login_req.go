package dto

type UserLoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
