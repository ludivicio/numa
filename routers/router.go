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

	// 后台首页
	beego.Router(adminURL, &admin.IndexController{}, "get:Index")

	// 后台登录
	beego.Router(adminURL+"/login", &admin.ProfileController{}, "*:Login")
	// 注销登录
	beego.Router(adminURL+"/logout", &admin.ProfileController{}, "*:Logout")
	// 查看修改管理员信息
	beego.Router(adminURL+"/profile", &admin.ProfileController{}, "*:Profile")
	// 上传头像
	beego.Router(adminURL+"/profile/head", &admin.ProfileController{}, "*:Head")
	// 修改登录密码
	beego.Router(adminURL+"/password", &admin.ProfileController{}, "*:Password")

	// 获取商品列表
	beego.Router(adminURL+"/item", &admin.ItemController{})
	// 获取商品分类列表
	beego.Router(adminURL+"/item/cate", &admin.ItemCateController{}, "*:List")
	// 添加商品分类
	beego.Router(adminURL+"/item/cate/add", &admin.ItemCateController{}, "*:Add")
	// 获取文章列表
	beego.Router(adminURL+"/article", &admin.ArticleController{})
	// 获取文章分类列表
	beego.Router(adminURL+"/article/cate", &admin.ArticleCateController{})
	// 获取商城列表
	beego.Router(adminURL+"/mall", &admin.MallController{})
	// 获取机器用户列表
	beego.Router(adminURL+"/autouser", &admin.AutoUserController{})
	// 获取数据信息，包含 爆料信息， 标签信息等
	beego.Router(adminURL+"/data", &admin.DataController{})
}
