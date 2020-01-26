package Repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

func getConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := 3306

	connectionString := dbUser + ":" + dbPass + "@" + dbHost + "/" + dbName + "?parseTime=true"
	println(connectionString)

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@("+dbHost+":"+strconv.Itoa(dbPort)+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
