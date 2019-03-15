package apollo

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

// Apollo apollo
type Apollo struct {
	ID         int    `orm:"column(lId)"`
	TraceID    string `orm:"column(strTraceId)"`
	ApolloType int    `orm:"column(nType)"`
	End        int    `orm:"column(nEnd)"`
}

// ApolloRuntimeData apollo运行时数据
type ApolloRuntimeData struct {
	ID         int    `orm:"column(lId)"`
	ApolloID   int    `orm:"column(lApolloId)"`
	TraceID    int    `orm:"column(strTraceId)"`
	BorrowMode int    `orm:"column(nBorrowMode)"`
	Type       int    `orm:"column(nType)"`
	NodeID     int    `orm:"column(lNodeId)"`
	NodeName   string `orm:"column(strNodeName)"`
	State      int    `orm:"column(nState)"`
	// Apollo      Apollo
}

type GraphLinkVO struct {
	From    int    `json:"from"`
	To      int    `json:"to"`
	Text    string `json:"text"`
	Visible bool   `json:"visible"`
}

type NodeVO struct {
	ApolloId   int    `json:"apolloId"`
	BorrowMode int    `json:"borrowMode"`
	Category   string `json:"category"`
	Color      string `json:"color"`
	Key        string `json:"key"`
	Loc        string `json:"loc"`
	Text       string `json:"text"`
	Figure     string `json:"figure"`
}

type GraphVO struct {
	LinkDataArray []GraphLinkVO `json:"linkDataArray"`
	NodeDataArray []NodeVO      `json:"nodeDataArray"`
}

func GetApolloBy(borrowMode, apolloType int, traceID string) Apollo {
	apollo := Apollo{}
	orm := orm.NewOrm()
	orm.Using("Apollo" + strconv.Itoa(borrowMode))
	orm.Raw("SELECT * FROM tbApolloData WHERE nBorrowMode = ? AND nType = ? AND strTraceId = ?", borrowMode, apolloType, traceID).QueryRow(&apollo)
	return apollo
}

// GetApolloRuntimeDataBy 获取ApolloRuntimeData
func GetApolloRuntimeDataBy(borrowMode, apolloType int, traceId string) ([]ApolloRuntimeData, error) {
	var datas []ApolloRuntimeData
	orm := orm.NewOrm()
	err := orm.Using("Apollo" + strconv.Itoa(borrowMode))
	if err != nil {
		return nil, err
	}
	orm.Raw("SELECT * FROM tbApolloRuntimeData WHERE strTraceId = ? AND nType = ?", traceId, apolloType).QueryRows(&datas)
	return datas, nil
}

// GetGraphVO 获取图结构
func GetGraphVO(borrowMode int, apolloType int, traceID string) GraphVO {
	var graph GraphVO
	nodeList := GetNodeListBy(borrowMode, apolloType)
	if nodeList == nil {
		return graph
	}
	relationList := GetNodeRelationBy(borrowMode, apolloType)
	if relationList == nil {
		return graph
	}
	runtimeDataList, _ := GetApolloRuntimeDataBy(borrowMode, apolloType, traceID)
	if runtimeDataList == nil {
		return graph
	}

	apolloID := runtimeDataList[0].ApolloID

	linkVOList := []GraphLinkVO{}
	for _, v := range relationList {
		linkVO := GraphLinkVO{}
		linkVO.From = v.FromNodeID
		linkVO.To = v.ToNodeID
		linkVO.Visible = true
		linkVO.Text = v.EdgeContent
		linkVOList = append(linkVOList, linkVO)
	}

	nodeVOList := []NodeVO{}
	for _, node := range nodeList {
		nodeVO := NodeVO{}
		nodeVO.ApolloId = apolloID
		nodeVO.Key = strconv.Itoa(node.ID)
		nodeVO.Loc = node.Location
		nodeVO.Text = node.NodeName
		nodeVO.BorrowMode = node.BorrowMode
		cat := node.Category
		if cat == 0 {
			nodeVO.Category = "Start"
		} else if cat == 1 {
			nodeVO.Category = "Common"
		} else if cat == 2 {
			nodeVO.Category = "Judge"
			nodeVO.Figure = "Diamond"
		} else if cat == 3 {
			nodeVO.Category = ""
		} else if cat == 4 {
			nodeVO.Category = "Wait"
			nodeVO.Figure = "Ellipse"
		} else {
			nodeVO.Category = ""
		}
		nodeVO.Color = getColorBy(node.ID, runtimeDataList)
		nodeVOList = append(nodeVOList, nodeVO)
	}

	graph.LinkDataArray = linkVOList
	graph.NodeDataArray = nodeVOList
	return graph
}

func getColorBy(nodeId int, runtimeDataList []ApolloRuntimeData) string {
	result := "LightSlateGray"
	for _, runtimeData := range runtimeDataList {
		if runtimeData.NodeID == nodeId {
			state := runtimeData.State
			if state == 0 {
				result = "LightSlateGray"
			} else if state == 1 {
				result = "LightSkyBlue"
			} else if state == 2 {
				result = "Turquoise"
			} else {
				result = "Red"
			}
		}
	}
	return result
}
