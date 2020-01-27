package Repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
)

var db *sql.DB

func GetConnection() {
	dbDriver := "mysql"
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := 3306

	var err error
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@("+dbHost+":"+strconv.Itoa(dbPort)+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return
}
