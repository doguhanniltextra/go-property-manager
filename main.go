package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/doguhanniltextra/property_go/database"
	"github.com/joho/godotenv"
)

type DatabaseC struct {
	DB *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Environment variables loaded from .env:")
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	// Şifreyi güvenlik için gösterme
	fmt.Println("DB_PASSWORD: [HIDDEN]")

	db, err := database.Connection()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	fmt.Println("Connection verified, now you can run queries.")

}
