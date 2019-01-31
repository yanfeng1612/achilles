package main

import (
	"achilles/engine/db"
	"achilles/engine/models"
	_ "achilles/engine/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	db.CreateDefaultDB("127.0.0.1:3306", "root", "root")
	models.Start()
	beego.Run()
}
