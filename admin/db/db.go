package db

import (
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

// DataBase 数据库基础数据结构
type DataBase struct {
	dbusername string
	dbpassowrd string
	dbhostsip  string
	dbname     string
	DB         *sqlx.DB
}

var databaseMap map[string]*DataBase

// Init 初始化DataBaseMap map
func Init() {
	databaseMap = make(map[string]*DataBase)

	username := "root"
	password := "root"

	// 初始Apollo配置自动化测试库
	apolloAutoDB := newDB("127.0.0.1:3306", "Achilles", username, password)
	databaseMap["Achilles"] = apolloAutoDB

	for _, v := range databaseMap {
		err := v.open()
		if err != nil {
			panic(err)
		}
	}
}

// GetDBByName 根据数据库名获取数据库操作API
func GetDBByName(name string) *DataBase {
	db := databaseMap[name]
	return db
}

func newDB(dbhostsip, dbname, dbusername, dbpassowrd string) *DataBase {
	return &DataBase{dbusername, dbpassowrd, dbhostsip, dbname, nil}
}

func (f *DataBase) open() error {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=30s&parseTime=True&multiStatements=false",
		f.dbusername, f.dbpassowrd, f.dbhostsip, f.dbname))

	db.SetMaxIdleConns(1024)
	db.SetConnMaxLifetime(time.Second * 3600)
	db.SetMaxOpenConns(1024)
	err = db.Ping()
	if err != nil {
		return err
	}
	f.DB = db
	return err
}

func (f *DataBase) close() { //关闭
	defer f.DB.Close()
}
