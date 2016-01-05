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
	beego.Router(adminURL + "/logout", &admin.IndexController{}, "get:Logout")
    // 后台首页
    beego.Router(adminURL, &admin.IndexController{}, "get:Index")
    // 修改个人信息
    beego.Router(adminURL + "/profile", &admin.IndexController{}, "get:Profile")
}
