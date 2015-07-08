package routers

import (
	"numa/handlers"
	"github.com/astaxie/beego"
)

func init() {

	// 前台首页
	beego.Router("/", &handlers.HomeHandler{})

	// 前台详情页
	beego.Router("/detail", &handlers.DetailHandler{})

	// 后台首页
    beego.Router("/admin", &handlers.MainHandler{})
}
