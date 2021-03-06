package dto

// RegisterDto is used when client post form /register url
type RegisterDto struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password.omitempty" form:"password.omitempty" validate:"ming:6"`
}
