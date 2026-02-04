package models

import (
	"github.com/jinzhu/gorm"
)

// User represents the user model in the database.
type User struct {
	// gorm.Model adds ID, CreatedAt, UpdatedAt, DeletedAt fields automatically
	gorm.Model

	// Name is required
	Name string `json:"name" binding:"required"`

	// Email must be unique and valid
	Email string `json:"email" gorm:"unique" binding:"required,email"`

	// Password is required (will be stored as a hash)
	Password string `json:"password" binding:"required"`
}
