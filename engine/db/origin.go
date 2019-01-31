package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DefaultDB *sqlx.DB
	DBName    = "Achilles"
	Username  = "writedafy"
	Password  = "writeDafy!@#$"
)

func CreateDefaultDB(host, username, passowrd string) *sqlx.DB {
	if DefaultDB != nil {
		return DefaultDB
	}
	db, err := sqlx.Connect("mysql", username+":"+passowrd+"@tcp("+host+")/"+DBName)
	if err != nil {
		panic(err)
	}
	DefaultDB = db
	err = DefaultDB.Ping()
	if err != nil {
		panic(err)
	}
	return DefaultDB
}

func Close() {
	if DefaultDB != nil {
		DefaultDB.Close()
	}
}
