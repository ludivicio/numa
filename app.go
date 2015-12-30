package main

import (
	_ "numa/routers"
	"github.com/astaxie/beego"
)

func main() {

	// 启用Session
	beego.SessionOn = true
	beego.SessionName = "numa"
	beego.AutoRender = true

	beego.CopyRequestBody = true

	beego.Run()
}

