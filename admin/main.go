package main

import (
	"github.com/astaxie/beego"
	"github.com/yanfeng1612/achilles/admin/db"
	"github.com/yanfeng1612/achilles/admin/models"
	_ "github.com/yanfeng1612/achilles/admin/routers"
	"time"
)

const (
	VERSION = "1.0.0"
)

func init() {
	var StartTime = time.Now().Unix()
	db.Init()
	models.Init(StartTime)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
	// libs.Main()
}
