package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type database struct {
	db *sql.DB
}

func NewDB() (*database, error) {

	var (
		dbuser  = os.Getenv("DBUSER")
		dbname  = os.Getenv("DBNAME")
		dbpass  = os.Getenv("DBPASS")
		connStr = fmt.Sprintf("user=%v dbname=%v password=%v port=5432 host=db sslmode=disable", dbuser, dbname, dbpass)
	)
	fmt.Printf("connStr %v \n", connStr)

	dbInstance, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := dbInstance.Ping(); err != nil {
		return nil, err
	}

	return &database{
		db: dbInstance,
	}, nil
}
