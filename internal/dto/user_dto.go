package dto

// UserResponse defines the payload returned to API clients.
type UserResponse struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

// UpdateUserRequest defines allowed profile updates.
type UpdateUserRequest struct {
    Name     string `json:"name" validate:"omitempty,min=3,max=100"`
    Email    string `json:"email" validate:"omitempty,email"`
    Password string `json:"password" validate:"omitempty,min=8,max=72"`
}
