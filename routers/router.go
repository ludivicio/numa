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

	var adminURL = beego.AppConfig.String("adminurl")
	
	// 后台登录
	beego.Router(adminURL + "/login", &admin.IndexController{}, "get:GoLogin;post:Login")
	// 注销登录
	beego.Router(adminURL + "/logout", &admin.IndexController{}, "post:Logout")
    // 后台首页
    beego.Router(adminURL , &admin.IndexController{}, "get:Index")
	
}
