package apollo

import (
	"github.com/yanfeng1612/achilles/admin/controllers"
	"github.com/yanfeng1612/achilles/admin/models/apollo"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type CaseController struct {
	controllers.BaseController
}

func (self *CaseController) List() {
	self.Data["pageTitle"] = "Case管理"
	self.Display()
}

func (self *CaseController) Add() {
	self.Data["pageTitle"] = "新增任务"
	self.Display()
}

func (self *CaseController) Edit() {
	self.Data["pageTitle"] = "编辑任务"
	id, _ := self.GetInt("id")
	apolloCase := apollo.GetApolloCaseById(id)
	self.Data["apolloCase"] = apolloCase
	self.Display()
}

func (self *CaseController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}
	result, totalCount := apollo.GetApolloCaseBy(page, limit)
	self.AjaxList("成功", MSG_OK, totalCount, result)
}

func (self *CaseController) Update() {
	caseName := self.GetString("caseName")
	inputParams := self.GetString("inputParams")
	expectResult := self.GetString("expectResult")
	id, _ := self.GetInt("id")
	apollo.UpdateApolloCaseBy(id, caseName, inputParams, expectResult)
	self.AjaxMsg("OK", MSG_OK)
}

func (this *CaseController) AjaxAdd() {
	apolloCase := apollo.ApolloCase{}
	apolloCase.CaseName = this.GetString("caseName")
	apolloCase.BorrowMode, _ = this.GetInt("borrowMode")
	apolloCase.ApolloType, _ = this.GetInt("apolloType")
	apolloCase.InputParams = this.GetString("inputParams")
	apolloCase.ExpectResult = this.GetString("expectResult")
	apollo.AddApolloCase(apolloCase)
	this.AjaxMsg("OK", MSG_OK)
}
