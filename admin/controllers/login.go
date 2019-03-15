package controllers

import (
	"fmt"
	"strconv"
	"time"

	"strings"

	"github.com/astaxie/beego"
	"github.com/yanfeng1612/achilles/admin/libs"
	"github.com/yanfeng1612/achilles/admin/models"
)

type LoginController struct {
	BaseController
}

func (self *LoginController) Login() {
	//if self.userId > 0 {
	//	self.redirect(beego.URLFor("HomeController.Index"))
	//}
	self.TplName = "login/login.html"
}

//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {

	//self.AjaxMsg("登录成功", MSG_OK)
	if self.userId > 0 {
		self.AjaxMsg("登录成功", MSG_OK)
	}

	if self.isPost() {
		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))
		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			fmt.Println(user)
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				self.AjaxMsg("帐号或密码错误", MSG_ERR)
			} else if user.Status == -1 {
				self.AjaxMsg("该帐号已禁用", MSG_ERR)
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				self.AjaxMsg("登录成功", MSG_OK)
			}
		}
	}
	self.AjaxMsg("请求方式错误", MSG_ERR)
}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("LoginController.Login"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}
