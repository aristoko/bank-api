package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() *sql.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	runMigration()
	return DB
}

func runMigration() {
	file, err := ioutil.ReadFile("internal/db/migrations/init.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = DB.Exec(string(file))
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully ✅")
}
