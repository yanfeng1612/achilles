package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yanfeng1612/achilles/admin/libs"
	"github.com/yanfeng1612/achilles/admin/models"
	"strconv"
	"strings"
	"time"
)

type TaskLogController struct {
	BaseController
}

func (self *TaskLogController) List() {
	taskId, err := self.GetInt("task_id")
	if err != nil {
		taskId = 1
	}

	task, err := models.TaskGetById(taskId)
	if err != nil {
		self.AjaxMsg(err.Error(), MSG_ERR)
	}
	self.Data["pageTitle"] = "日志管理 - " + task.TaskName + "(#" + strconv.Itoa(task.Id) + ")"
	self.Data["task_id"] = task.Id
	self.Display()
}

func (self *TaskLogController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	taskId, err := self.GetInt("task_id")
	if err != nil {
		taskId = 1
	}

	TextStatus := []string{
		"<font color='orange'><i class='fa fa-question-circle'></i> 超时</font>",
		"<font color='red'><i class='fa fa-times-circle'></i> 错误</font>",
		"<font color='green'><i class='fa fa-check-square'></i> 正常</font>",
	}

	Status, err := self.GetInt("status")

	if err == nil && Status != 9 {
		filters = append(filters, "status", Status)
	}
	filters = append(filters, "task_id", taskId)

	result, count := models.TaskLogGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["task_id"] = v.TaskId
		row["start_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["process_time"] = float64(v.ProcessTime) / 1000
		if v.Status == 0 {
			row["output_size"] = libs.SizeFormat(float64(len(v.Output)))
		} else {
			row["output_size"] = libs.SizeFormat(float64(len(v.Error)))
		}
		index := v.Status + 2
		if index > 2 {
			index = 2
		}
		row["status"] = TextStatus[index]

		list[k] = row
	}

	self.AjaxList("成功", MSG_OK, count, list)
}

func (self *TaskLogController) Detail() {

	//日志内容
	id, _ := self.GetInt("id")
	tasklog, err := models.TaskLogGetById(id)
	if err != nil {
		self.Ctx.WriteString("日志不存在")
		return
	}
	LogTextStatus := []string{
		"<font color='orange'><i class='fa fa-question-circle'></i>超时</font>",
		"<font color='red'><i class='fa fa-times-circle'></i> 错误</font>",
		"<font color='green'><i class='fa fa-check-square'></i> 正常</font>",
	}
	row := make(map[string]interface{})
	row["id"] = tasklog.Id
	row["task_id"] = tasklog.TaskId
	row["start_time"] = beego.Date(time.Unix(tasklog.CreateTime, 0), "Y-m-d H:i:s")
	row["process_time"] = float64(tasklog.ProcessTime) / 1000
	if tasklog.Status == 0 {
		row["output_size"] = libs.SizeFormat(float64(len(tasklog.Output)))
	} else {
		row["output_size"] = libs.SizeFormat(float64(len(tasklog.Error)))
	}

	row["output"] = tasklog.Output
	row["error"] = tasklog.Error

	index := tasklog.Status + 2
	if index > 2 {
		index = 2
	}
	row["status"] = LogTextStatus[index]

	self.Data["taskLog"] = row

	//任务详情
	task, err := models.TaskGetById(tasklog.TaskId)
	if err != nil {
		self.AjaxMsg(err.Error(), MSG_ERR)
	}
	TextStatus := []string{
		"<font color='red'><i class='fa fa-minus-square'></i> 暂停</font>",
		"<font color='green'><i class='fa fa-check-square'></i> 运行中</font>",
		"<font color='orange'><i class='fa fa-check-square'></i> 待审核</font>",
		"<font color='blue'><i class='fa fa-times-circle'></i> 审核失败</font>",
	}

	self.Data["TextStatus"] = TextStatus[task.Status]
	self.Data["CreateTime"] = beego.Date(time.Unix(task.CreateTime, 0), "Y-m-d H:i:s")
	self.Data["UpdateTime"] = beego.Date(time.Unix(task.UpdateTime, 0), "Y-m-d H:i:s")
	self.Data["task"] = task
	// 分组列表
	self.Data["taskGroup"] = taskGroupLists(self.taskGroups, self.userId)

	serverName := "本地服务器"
	if task.ServerId == 0 {
		serverName = "本地服务器"
	} else {
		server, err := models.TaskServerGetById(task.ServerId)
		if err == nil {
			serverName = server.ServerName
		}
	}
	self.Data["serverName"] = serverName

	//任务分组
	groupName := "默认分组"
	if task.GroupId > 0 {
		group, err := models.TaskGroupGetById(task.GroupId)
		if err == nil {
			groupName = group.GroupName
		}
	}
	self.Data["GroupName"] = groupName

	//创建人和修改人
	createName := "未知"
	updateName := "未知"
	if task.CreateId > 0 {
		admin, err := models.AdminGetById(task.CreateId)
		if err == nil {
			createName = admin.RealName
		}
	}

	if task.UpdateId > 0 {
		admin, err := models.AdminGetById(task.UpdateId)
		if err == nil {
			updateName = admin.RealName
		}
	}

	//是否出错通知
	self.Data["adminInfo"] = []int{0}
	if task.NotifyUserIds != "0" && task.NotifyUserIds != "" {
		self.Data["adminInfo"] = AllAdminInfo(task.NotifyUserIds)
	}
	self.Data["CreateName"] = createName
	self.Data["UpdateName"] = updateName
	self.Data["pageTitle"] = "日志详细" + "(#" + strconv.Itoa(id) + ")"
	self.Display()
}

// 批量操作日志
func (self *TaskLogController) AjaxDel() {
	ids := self.GetStrings("ids")
	idArr := strings.Split(ids[0], ",")

	if len(idArr) < 1 {
		self.AjaxMsg("请选择要操作的项目", MSG_ERR)
	}

	for _, v := range idArr {
		id, _ := strconv.Atoi(v)
		if id < 1 {
			continue
		}
		models.TaskLogDelById(id)
	}

	self.AjaxMsg("", MSG_OK)
}
