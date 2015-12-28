package routers

import (
	"numa/handlers/admin"
	"numa/handlers/home"
	"github.com/astaxie/beego"
)

func init() {

	// 前台首页
	beego.Router("/", &home.HomeController{}, "*:Index")

	// 前台详情页
	beego.Router("/detail", &home.DetailController{})

	// 后台首页
    beego.Router("/admin", &admin.MainController{})
}
