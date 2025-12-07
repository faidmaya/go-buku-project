package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded; relying on environment variables")
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	if ssl == "" {
		ssl = "require"
	}

	// Convert port ke int
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("database open error: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database ping error: %v", err)
	}

	log.Println("Successfully connected to the database!")
	DB = db
}
