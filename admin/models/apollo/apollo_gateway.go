package apollo

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type GatewayRecord struct {
	FundID       int    `orm:"column(nFundId)"`
	SN           string `orm:"column(strSN)"`
	ApolloID     int    `orm:"column(lApolloId)"`
	NodeID       int    `orm:"column(lNodeId)"`
	RequestData  string `orm:"column(strRequestData)"`
	ResponseData string `orm:"column(strResponseData)"`
	CreateTime   string `orm:"column(dtCreateTime)"`
}

func GetGatewayRecorListBy(fundID int, apolloID int, nodeID int) []GatewayRecord {
	orm := orm.NewOrm()
	orm.Using("Apollo" + strconv.Itoa(fundID))
	var list []GatewayRecord
	orm.Raw("SELECT * FROM tbApolloGatewayRecord WHERE nFundId = ? AND lApolloId = ? AND lNodeId = ? ORDER BY lId DESC", fundID, apolloID, nodeID).QueryRows(&list)
	return list
}
