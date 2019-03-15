package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

type Database struct {
	Name      string
	TableList []Table
}

type Table struct {
	DdatabaseName string
	Name          string
	ColumnList    []Column
}

type Column struct {
	TableSchema   string `orm:"column(TABLE_SCHEMA)"`
	TableName     string `orm:"column(TABLE_NAME)"`
	ColumnName    string `orm:"column(COLUMN_NAME)"`
	ColumnDefault string `orm:"column(COLUMN_DEFAULT)"`
	IsNullable    bool   `orm:"column(IS_NULLABLE)"`
	DataType      string `orm:"column(DATA_TYPE)"`
	Length        int    `orm:"column(CHARACTER_OCTET_LENGTH)"`
	Comment       string `orm:"column(COLUMN_COMMENT)"`
}

func GetTableInfo(databaseName string, tableName string) *Table {
	table := &Table{
		DdatabaseName: databaseName,
		Name:          tableName,
	}
	orm := orm.NewOrm()
	err := orm.Using("information_schema")
	if err != nil {
		return table
	}
	columns := make([]Column, 0)
	_, err = orm.Raw("SELECT TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,IS_NULLABLE,COLUMN_DEFAULT,DATA_TYPE,CHARACTER_OCTET_LENGTH,COLUMN_COMMENT FROM columns WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?", databaseName, tableName).QueryRows(&columns)
	if err != nil {
		return table
	}
	table.ColumnList = columns
	return table
}

func PrintfJDBCTemplate(table *Table) {
	result := "INSERT INTO '" + table.DdatabaseName + "'.'" + table.Name + "' ("
	q := ""
	for i, column := range table.ColumnList {
		result += column.ColumnName
		q += "?"
		if i != len(table.ColumnList)-1 {
			result += ","
			q += ","
		}
	}
	result += ") VALUES ("
	result += q
	result = result + ");"
	log.Printf(result)
}
