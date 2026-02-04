package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold the database connection.
// It is used by other packages to interact with the database.
var DB *gorm.DB

// Connect establishes a connection to the PostgreSQL database using GORM.
// It reads connection details from environment variables.
func Connect() {
	// Read environment variables loaded from .env (or system env)
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Create the Data Source Name (DSN) string required by the Postgres driver.
	// sslmode=disable is used for development/local environments.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	// Open the connection using GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails, log the error and stop the application
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Connected to database")
}
