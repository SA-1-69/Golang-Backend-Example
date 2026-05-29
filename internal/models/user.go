package models

import "gorm.io/gorm"

// User represents a user account in the system. It embeds gorm.Model
// to include ID, CreatedAt, UpdatedAt and DeletedAt fields.
type User struct {
    gorm.Model
    Name     string `gorm:"size:100;not null" json:"name"`
    Email    string `gorm:"size:100;uniqueIndex;not null" json:"email"`
    Password string `gorm:"size:255;not null" json:"-"`
}
