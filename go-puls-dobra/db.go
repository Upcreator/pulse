// db.go
package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func setupDB() *sql.DB {
	// Подключение к конкрентной БД
	connStr := "user=postgres password=mysecretpassword dbname=postgres host=87.228.19.168 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Отладочное сообщение для проверки подключения к БД
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
