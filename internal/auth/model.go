package auth

type ErrorResponse struct {
	Message string `json:"message"`
}

type SignUpRequest struct {
	Name            string `form:"name" binding:"required"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirmPassword" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
