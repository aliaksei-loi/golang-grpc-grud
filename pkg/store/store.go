package store

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "grpcdb"
)

// Connect to postgres db
func Connect() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if user == "" || password == "" {
		log.Fatal("DB_USER and DB_PASSWORD environment variables must be set")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
