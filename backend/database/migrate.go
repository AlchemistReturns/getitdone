package database

import (
	"fmt"

	"example.com/getitdone/internal/models"
)

func Migrate() {
	fmt.Println("Migrating database...")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}
