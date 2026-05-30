package dto

import "time"

// UserResponse defines the payload returned to API clients.
type UserResponse struct {
	ID        uint       `json:"id" example:"1"`
	FirstName string     `json:"first_name" example:"John"`
	LastName  string     `json:"last_name" example:"Doe"`
	Email     string     `json:"email" example:"john.doe@example.com"`
	Age       int        `json:"age" example:"30"`
	BirthDay  *time.Time `json:"birth_day" example:"1993-05-30T00:00:00Z"`
	GenderID  *uint      `json:"gender_id" example:"1"`
	CreatedAt string     `json:"created_at" example:"2025-05-30T10:30:00Z"`
	UpdatedAt string     `json:"updated_at" example:"2025-05-30T10:30:00Z"`
}

// UpdateUserRequest defines allowed profile updates.
type UpdateUserRequest struct {
	FirstName string     `json:"first_name" validate:"omitempty,min=2,max=100" example:"John"`
	LastName  string     `json:"last_name" validate:"omitempty,min=2,max=100" example:"Doe"`
	Email     string     `json:"email" validate:"omitempty,email" example:"john.doe@example.com"`
	Password  string     `json:"password" validate:"omitempty,min=8,max=72" example:"newpassword123"`
	Age       *int       `json:"age" validate:"omitempty,min=0,max=150" example:"31"`
	BirthDay  *time.Time `json:"birth_day" example:"1993-05-30T00:00:00Z"`
	GenderID  *uint      `json:"gender_id" example:"1"`
}
