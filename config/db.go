package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// function to open connection to mysql database
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "facebookdb"
	dbIP := "127.0.0.1:3306"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db

}