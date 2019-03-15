package apollo

import (
	"github.com/jmoiron/sqlx"
	"github.com/yanfeng1612/achilles/admin/db"
	"log"
)

// Node 节点
type Node struct {
	ID         int    `db:"lId"`
	NodeName   string `db:"strNodeName"`
	BorrowMode int    `db:"nBorrowMode"`
	ApolloType int    `db:"nApolloType"`
	Location   string `db:"strLocation"`
	Category   int    `db:"nCategory"` //节点分类 0 - 开始节点 1 - 普通节点 2 - 判断节点 3 - 结束节点
}

// NodeRelation 节点关系
type NodeRelation struct {
	FromNodeID  int    `db:"lFromNodeId"`
	ToNodeID    int    `db:"lToNodeId"`
	EdgeContent string `db:"strEdgeContent"`
}

// GetNodeListBy 根据条件获取节点列表
func GetNodeListBy(borrowMode, apolloType int) []Node {
	var list []Node
	sql := "SELECT lId,strNodeName,nBorrowMode,nApolloType,strLocation,nCategory FROM tbNode WHERE nBorrowMode =:nBorrowMode AND nApolloType =:nApolloType"
	var rows *sqlx.Rows
	arg := map[string]interface{}{"nBorrowMode": borrowMode, "nApolloType": apolloType}

	db := db.GetDBByName("Apollo")

	rows, _ = db.DB.NamedQuery(sql, arg)

	err := sqlx.StructScan(rows, &list)
	if err != nil {
		log.Println(err)
	}
	return list
}

// GetNodeRelationBy 根据条件获取节点关系列表
func GetNodeRelationBy(borrowMode, apolloType int) []NodeRelation {
	var list []NodeRelation
	sql := "SELECT lFromNodeId,lToNodeId,strEdgeContent FROM tbNodeRelation WHERE nBorrowMode =:nBorrowMode AND nApolloType =:nApolloType"
	var rows *sqlx.Rows
	arg := map[string]interface{}{"nBorrowMode": borrowMode, "nApolloType": apolloType}

	rows, _ = db.GetDBByName("Apollo").DB.NamedQuery(sql, arg)

	err := sqlx.StructScan(rows, &list)
	if err != nil {
		log.Println(err)
	}
	return list
}
