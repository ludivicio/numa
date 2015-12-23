package routers

import (
	"numa/handlers/admin"
	"numa/handlers/home"
	"github.com/astaxie/beego"
)

func init() {

	// 前台首页
	beego.Router("/", &home.HomeHandler{}, "*:Index")

	// 前台详情页
	beego.Router("/detail", &home.DetailHandler{})

	// 后台首页
    beego.Router("/admin", &admin.MainHandler{})
}
