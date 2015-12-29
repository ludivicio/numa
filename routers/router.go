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

	var adminURL = "/admin"
	
	// 后台首页
    beego.Router(adminURL , &admin.MainController{})
	
	// 后台登录
	beego.Router(adminURL + "/login", &admin.LoginController{}, "get:Index;post:Login")
	
	
}
