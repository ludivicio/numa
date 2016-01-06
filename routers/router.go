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
	beego.Router(adminURL + "/login", &admin.IndexController{}, "*:Login")
	// 注销登录
	beego.Router(adminURL + "/logout", &admin.IndexController{}, "*:Logout")
    // 后台首页
    beego.Router(adminURL, &admin.IndexController{}, "get:Index")
    // 查看个人信息
    beego.Router(adminURL + "/profile", &admin.IndexController{}, "*:Profile")
	// 修改登录密码
	beego.Router(adminURL + "/password", &admin.IndexController{}, "*:Password")
}
