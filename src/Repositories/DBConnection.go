package Repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB


func GetConnection() {
	dbDriver := "postgres"
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	schema := os.Getenv("SCHEMA")
	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema)


	var err error
	db, err = sql.Open(dbDriver, psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return
}
