package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

var (
	Username = "writedafy"
	Password = "writeDafy!@#$"
)

func CreateDB(host, dbName, username, passowrd string) *sql.DB {
	db, err := sql.Open("mysql", username+":"+passowrd+"@tcp("+host+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateDBWithUser(host, dbName string) *sql.DB {
	return CreateDB(host, dbName, Username, Password)
}
