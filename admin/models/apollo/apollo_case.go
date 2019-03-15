package apollo

import (
	"github.com/astaxie/beego/orm"
	"log"
)

// ApolloCase Case
type ApolloCase struct {
	ID           int    `orm:"column(lId)"`
	CaseName     string `orm:"column(strCaseName)"`
	BorrowMode   int    `orm:"column(nBorrowMode)"`
	ApolloType   int    `orm:"column(nApolloType)"`
	InputParams  string `orm:"column(strInputParams)"`
	ExpectResult string `orm:"column(strExpectResult)"`
}

// GetRunableApolloCaseBy 获取待执行的ApolloCase
func GetApolloCaseBy(page, pageSize int) ([]ApolloCase, int64) {
	offset := (page - 1) * pageSize
	var list []ApolloCase
	orm := orm.NewOrm()
	var totalCount int64
	orm.Raw("SELECT COUNT(*) AS count FROM tbApolloCase").QueryRow(&totalCount)
	_, err := orm.Raw("SELECT * FROM tbApolloCase ORDER BY lId DESC LIMIT ?,?", offset, pageSize).QueryRows(&list)
	if err != nil {
		log.Println(err)
	}
	return list, totalCount
}

// GetRunableApolloCaseBy 获取ApolloCase slice
func GetApolloCaseById(id int) ApolloCase {
	var apolloCase ApolloCase
	orm := orm.NewOrm()
	orm.Using("Apollo_Auto")
	err := orm.Raw("SELECT * FROM tbApolloCase WHERE lId = ?", id).QueryRow(&apolloCase)
	if err != nil {
		log.Println(err)
	}
	return apolloCase
}

func UpdateApolloCaseBy(id int, caseName, inputParams, expectResult string) {
	orm := orm.NewOrm()
	orm.Raw("UPDATE tbApolloCase SET strCaseName = ?, strInputParams = ?,strExpectResult = ? WHERE lId = ?", caseName, inputParams, expectResult, id).Exec()
}

func AddApolloCase(apolloCase ApolloCase) {
	orm.NewOrm().Raw("INSERT INTO tbApolloCase (strCaseName,nBorrowMode,nApolloType,dtScheduleTime,strInputParams,strExpectResult,dtCreateTime) VALUES (?,?,?,NOW(),?,?,NOW())", apolloCase.CaseName, apolloCase.BorrowMode, apolloCase.ApolloType, apolloCase.InputParams, apolloCase.ExpectResult).Exec()
}
