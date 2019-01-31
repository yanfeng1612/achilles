package main

import (
	"github.com/astaxie/beego"
	"github.com/yanfeng1612/achilles/engine/db"
	"github.com/yanfeng1612/achilles/engine/models"
	_ "github.com/yanfeng1612/achilles/engine/routers"
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
