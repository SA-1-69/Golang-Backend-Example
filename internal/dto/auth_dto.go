package dto

// RegisterRequest contains the fields needed to create a new user.
type RegisterRequest struct {
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,max=72"`
}

// LoginRequest contains the fields needed for user login.
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

// AuthResponse contains the JWT token and user profile returned after auth.
type AuthResponse struct {
    Token string       `json:"token"`
    User  UserResponse `json:"user"`
}
