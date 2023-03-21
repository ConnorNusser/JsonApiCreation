package main

import (
	"database/sql"
	"fmt"
)

func main() {
	store, err := dbCreation()
	if err != nil {
		return
	}
	store.Init()
	server := NewApiServer(":8080", store)
	server.Run()
	fmt.Println("Hello")
	dbCreation()
}

type postgresStore struct {
	db *sql.DB
}

func dbCreation() (*postgresStore, error) {
	connStr := "user=connor dbName=postgres password=hi sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if pingError := db.Ping(); pingError != nil {
		return nil, pingError
	}
	return &postgresStore{
		db: db,
	}, nil

}
