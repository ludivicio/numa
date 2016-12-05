package main

import (
	_ "numa/routers"
	"github.com/astaxie/beego"
)

func main() {

	// 启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "numa"
	beego.BConfig.WebConfig.AutoRender = true

	beego.BConfig.CopyRequestBody = true

	beego.Run()
}

