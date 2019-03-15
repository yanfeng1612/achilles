package controllers

import (
	"github.com/yanfeng1612/achilles/admin/libs"
	"github.com/yanfeng1612/achilles/admin/models"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (self *UserController) Edit() {
	self.Data["pageTitle"] = "资料修改"
	id := self.userId
	Admin, _ := models.AdminGetById(id)
	row := make(map[string]interface{})
	row["id"] = Admin.Id
	row["login_name"] = Admin.LoginName
	row["real_name"] = Admin.RealName
	row["phone"] = Admin.Phone
	row["email"] = Admin.Email
	self.Data["admin"] = row
	self.Display()
}

func (self *UserController) AjaxSave() {
	Admin_id, _ := self.GetInt("id")
	Admin, _ := models.AdminGetById(Admin_id)
	//修改
	Admin.Id = Admin_id
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = self.userId
	Admin.LoginName = strings.TrimSpace(self.GetString("login_name"))
	Admin.RealName = strings.TrimSpace(self.GetString("real_name"))
	Admin.Phone = strings.TrimSpace(self.GetString("phone"))
	Admin.Email = strings.TrimSpace(self.GetString("email"))

	resetPwd := self.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(self.GetString("password_old"))
		pwdOldMd5 := libs.Md5([]byte(pwdOld + Admin.Salt))
		if Admin.Password != pwdOldMd5 {
			self.AjaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(self.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(self.GetString("password_new2"))

		if len(pwdNew1) < 6 {
			self.AjaxMsg("密码长度需要六位以上", MSG_ERR)
		}
		if pwdNew1 != pwdNew2 {
			self.AjaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := libs.Password(4, pwdNew1)
		Admin.Password = pwd
		Admin.Salt = salt
	}
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = self.userId
	Admin.Status = 1

	if err := Admin.Update(); err != nil {
		self.AjaxMsg(err.Error(), MSG_ERR)
	}
	self.AjaxMsg("", MSG_OK)
}
