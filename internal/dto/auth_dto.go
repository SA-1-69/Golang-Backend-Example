package dto

// RegisterRequest contains the fields needed to create a new user.
type RegisterRequest struct {
	FirstName string `json:"first_name" validate:"required,min=2,max=100" example:"John"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100" example:"Doe"`
	Email     string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password  string `json:"password" validate:"required,min=8,max=72" example:"password123"`
	Age       int    `json:"age" validate:"omitempty,min=0,max=150" example:"30"`
	GenderID  *uint  `json:"gender_id" validate:"omitempty" example:"1"`
}

// LoginRequest contains the fields needed for user login.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

// AuthResponse contains the JWT token and user profile returned after auth.
type AuthResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}
