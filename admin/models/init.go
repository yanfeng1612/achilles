package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yanfeng1612/achilles/admin/libs"
	"github.com/yanfeng1612/achilles/admin/models/apollo"
)

var StartTime int64

func Init(startTime int64) {
	StartTime = startTime
	orm.RegisterDataBase("default", "mysql", "writedafy:writeDafy!@#$@tcp(127.0.0.1:3306)/Achilles?charset=utf8")

	orm.RegisterModel(
		new(Admin),
		new(Auth),
		new(Role),
		new(RoleAuth),
		new(ServerGroup),
		new(TaskServer),
		new(Ban),
		new(Group),
		new(Task),
		new(TaskLog),
		new(apollo.ApolloRuntimeData),
		new(apollo.ApolloCase),
	)

	libs.InitRedisClient("192.168.0.128:6379")

	// initRedisInfo()

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return "pp_" + name
}
