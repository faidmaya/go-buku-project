package main

import (
	"log"
	"os"

	"go-buku-project/database"
	"go-buku-project/routers"

	"github.com/joho/godotenv"
)

func main() {
	// load env from .env if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded; relying on environment variables")
	}
	database.Connect()

	useJWT := true // set false to use Basic Auth instead

	r := routers.SetupRouter(useJWT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
