package main

import (
	_ "numa/routers"
	"github.com/astaxie/beego"
)

func main() {

	beego.SessionOn = true
	beego.SessionName = "numa"
	beego.AutoRender = true

	beego.CopyRequestBody = true

	beego.Run()
}

