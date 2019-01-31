package models

import (
	"fmt"
	"github.com/yanfeng1612/achilles/engine/db"
)

type Graph struct {
	Id               int64  `db:"id"`
	Name             string `db:"name"`
	CreateTime       string `db:"createTime"`
	NodeList         []Node
	NodeRelationList []NodeRelation
}

// Node 节点
type Node struct {
	Id              int64  `db:"id"`
	GraphId         int64  `db:"graphId"`
	NodeName        string `db:"name"`
	Location        string `db:"location"`
	Type            int    `db:"type"` //节点分类 0 - 开始节点 1 - 普通节点 2 - 判断节点 3 - 结束节点
	ExecType        int    `db:"execType"`
	Priority        int    `db:"priority"`
	InitDelay       int    `db:"initDelay"`
	RetryNumber     int    `db:"retryNumber"`
	Url             string `db:"url"`
	ReqParam        string `db:"reqParam"`
	ExecSql         string `db:"execSql"`
	CreateTime      string `db:"createTime"`
	RefreshTime     string `db:"refreshTime"`
	HttpContentType string
}

// NodeRelation 节点关系
type NodeRelation struct {
	Id            int64  `db:"id"`
	GraphId       int64  `db:"graphId"`
	FromNodeID    int64  `db:"fromNodeId"`
	ToNodeID      int64  `db:"toNodeId"`
	EdgnCondition string `db:"edgnCondition"`
	CreateTime    string `db:"createTime"`
	RefreshTime   string `db:"refreshTime"`
	FromNode      Node
	ToNode        Node
}

func getGraphBy(id int64) (*Graph, error) {
	graph := &Graph{}
	db.DefaultDB.Get(graph, "SELECT id,name FROM graph WHERE id = ?", id)

	nodeList := []Node{}
	err := db.DefaultDB.Select(&nodeList, "SELECT * FROM node WHERE graphId = ?", id)
	if err != nil {
		fmt.Println(err)
		return graph, err
	}
	graph.NodeList = nodeList

	relationList := []NodeRelation{}
	err = db.DefaultDB.Select(&relationList, "SELECT * FROM node_relation WHERE graphId = ?", id)
	if err != nil {
		fmt.Println(err)
		return graph, err
	}
	graph.NodeRelationList = relationList

	for _, relation := range graph.NodeRelationList {
		nodeMap := getMapFromAndToNode(graph.NodeList, relation.FromNodeID, relation.ToNodeID)
		relation.FromNode = nodeMap[relation.FromNodeID]
		relation.ToNode = nodeMap[relation.ToNodeID]
	}

	return graph, nil
}

func getMapFromAndToNode(nodes []Node, fromNodeId, toNodeId int64) map[int64]Node {
	m := make(map[int64]Node)
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if node.Id == fromNodeId {
			m[fromNodeId] = node
			continue
		}
		if node.Id == toNodeId {
			m[toNodeId] = node
		}
	}
	return m
}

func getNextNode(graph *Graph, current *Node) []Node {
	result := []Node{}
	for _, relation := range graph.NodeRelationList {
		if relation.FromNodeID == current.Id {
			result = append(result, relation.ToNode)
		}
	}
	return result
}
