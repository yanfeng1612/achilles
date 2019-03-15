package routers

import (
	"github.com/astaxie/beego"
	"github.com/yanfeng1612/achilles/admin/controllers"
	"github.com/yanfeng1612/achilles/admin/controllers/apollo"
)

func init() {

	// 默认登录
	beego.Router("/", &controllers.LoginController{}, "*:Login")
	beego.Router("/login_in", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/help", &controllers.HomeController{}, "*:Help")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.TaskLogController{})

	//资源分组管理
	beego.AutoRouter(&controllers.ServerGroupController{})
	beego.AutoRouter(&controllers.ServerController{})
	beego.AutoRouter(&controllers.BanController{})

	//权限用户相关
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

	//apollo相关
	beego.AutoRouter(&apollo.ApolloJobController{})

	beego.AutoRouter(&controllers.PerfmanceController{})

}
