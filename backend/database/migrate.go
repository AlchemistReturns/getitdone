package database

import (
	"fmt"

	"example.com/getitdone/internal/models"
)

// Migrate automatically updates the database schema to match the code's data models.
// It creates tables, missing columns, and missing indexes.
// WARNING: It will NOT delete unused columns to protect data.
func Migrate() {
	fmt.Println("Migrating database...")

	// AutoMigrate will create the "users" table based on the User struct.
	// You can add more models here as comma-separated arguments.
	DB.AutoMigrate(&models.User{})

	fmt.Println("Database migrated")
}
