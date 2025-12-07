package main

import (
	"fmt"
	"log"
	"os"

	"go-buku-project/database"
	"go-buku-project/routers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables dari .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded; relying on environment variables")
	}

	// Debug: cek env variable
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_SSLMODE:", os.Getenv("DB_SSLMODE"))

	// Connect ke database Railway
	database.Connect()

	// Gunakan JWT untuk autentikasi
	useJWT := true // set false untuk Basic Auth

	// Setup router Gin
	r := routers.SetupRouter(useJWT)

	// Ambil port dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default 8080
	}

	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
