package apollo

import (
	"encoding/json"
	"github.com/yanfeng1612/achilles/admin/controllers"
	"github.com/yanfeng1612/achilles/admin/models/apollo"
)

type ApolloJobController struct {
	controllers.BaseController
}

func (self *ApolloJobController) JobList() {
	self.Data["pageTitle"] = "任务管理"
	self.Display()
}

func (this *ApolloJobController) Detail() {
	this.Data["pageTitle"] = "任务详情"
	this.Data["borrowMode"], _ = this.GetInt("borrowMode")
	this.Data["apolloType"], _ = this.GetInt("apolloType")
	this.Data["traceId"] = this.GetString("traceId")
	this.Display()
}

func (this *ApolloJobController) Gateway() {
	this.Data["pageTitle"] = "网关报文"
	fundID, _ := this.GetInt("borrowMode")
	apolloID, _ := this.GetInt("apolloId")
	nodeID, _ := this.GetInt("nodeId")
	list := apollo.GetGatewayRecorListBy(fundID, apolloID, nodeID)
	this.Data["list"] = list
	this.Display()
}

func (self *ApolloJobController) GetApolloRuntimeData() {
	borrowMode, _ := self.GetInt("borrowMode")
	traceID := self.GetString("traceId")
	nType := 1
	graph := apollo.GetGraphVO(borrowMode, nType, traceID)
	jsonString, _ := json.Marshal(graph)
	self.Data["json"] = string(jsonString)
	self.ServeJSON()
	self.StopRun()
}
