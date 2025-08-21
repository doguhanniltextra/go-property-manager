package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lib/pq"
)

var DB *sql.DB

func Connection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	cfg, err := pq.NewConnector(dsn)
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(cfg)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")
	return db, nil
}
