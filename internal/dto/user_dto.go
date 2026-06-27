package dto

import "time"

type UserResponse struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	BirthDay  *time.Time `json:"birth_day"`
	GenderID  *uint      `json:"gender_id"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

type UpdateUserRequest struct {
	FirstName string     `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName  string     `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email     string     `json:"email" validate:"omitempty,email"`
	Password  string     `json:"password" validate:"omitempty,min=8,max=72"`
	Age       *int       `json:"age" validate:"omitempty,min=0,max=150"`
	BirthDay  *time.Time `json:"birth_day"`
	GenderID  *uint      `json:"gender_id"`
}
