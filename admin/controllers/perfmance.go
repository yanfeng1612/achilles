package controllers

import (
	// "log"
	"github.com/yanfeng1612/achilles/admin/models"
)

type PerfmanceController struct {
	BaseController
}

func (this *PerfmanceController) Main() {
	this.Data["pageTitle"] = "性能测试"
	this.Display()
}

func (this *PerfmanceController) Run() {
	url := this.GetString("url")
	para := this.GetString("para")
	concurrentNum, _ := this.GetInt("concurrentNum")
	requestNum, _ := this.GetInt("requestNum")
	request := &models.PermanceRequest{
		Url:           url,
		Para:          para,
		ConcurrentNum: concurrentNum,
		RequestNum:    requestNum,
	}
	models.PerfermanceGo(*request)
	this.AjaxMsg("OK", MSG_OK)
}
