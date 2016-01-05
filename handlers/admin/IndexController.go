package admin

import (
	"encoding/base64"
	"fmt"
	"numa/models"
	"numa/toolkit"
	"strings"

	"github.com/astaxie/beego"
)

// IndexController 后台首页控制器
type IndexController struct {
	BaseController
}

// Index 后台首页
func (m *IndexController) Index() {
	m.Data["title"] = "首页"
	m.display()
}

// GoLogin 跳转至登录页面
func (m *IndexController) GoLogin() {
	m.TplNames = "admin/login.html"
}

// Login 登录处理
func (m *IndexController) Login() {

	account := strings.TrimSpace(m.GetString("account"))
	password := strings.TrimSpace(m.GetString("password"))

	fmt.Println("account = " + account)
	fmt.Println("password = " + password)

	remember := m.GetString("remember")

	fmt.Println(toolkit.GenUID())

	if account != "" && password != "" {
		var admin models.Admin
		admin.Account = account

		// if account == "admin" && password == "123456" {
		// 	admin.LastIP = m.GetClientIP()
		// 	admin.Password = toolkit.SHA256([]byte(password))
		//     admin.LastTime = toolkit.GetTime()
		// 	admin.Token = toolkit.GenUID()
		// 	admin.Email = "lurma@qq.com"
		// 	admin.Status = 1
		// 	admin.Insert()
		// 	m.Redirect(beego.AppConfig.String("adminurl"), 302)
		// 	return
		// }

		if admin.Read("account") != nil || admin.Password != toolkit.SHA256([]byte(password)) {
			// 用户名或密码输入错误
			fmt.Println("账号或密码填写错误")
			m.Data["errmsg"] = "账号或密码填写错误"
		} else {
			// 登录成功
			token := toolkit.GenUID()

			fmt.Println("client ip = " + m.GetClientIP())
			
			admin.LastIP = m.GetClientIP()
			admin.LastTime = toolkit.GetTime()
			admin.Token = token
			admin.Update()

			fmt.Println("key = " + beego.AppConfig.String("aeskey"))
			key := []byte(beego.AppConfig.String("aeskey"))
			result, err := toolkit.AesEncrypt([]byte(m.GetClientIP()+"|"+token), key)
			if err != nil {
				return
			}

			auth := base64.StdEncoding.EncodeToString(result)
			fmt.Println("auth = " + auth)

			if remember == "yes" {
				m.Ctx.SetCookie("auth", auth, 7*86400)
			} else {
				m.Ctx.SetCookie("auth", auth)
			}

			m.Redirect(beego.AppConfig.String("adminurl"), 302)
		}
	} else {
		// 用户名或密码为空
		fmt.Println("账号或密码为空")
		m.Data["errmsg"] = "账号或密码不能为空"
	}

	m.TplNames = "admin/login.html"
}

// Logout 注销登录
func (m *IndexController) Logout() {
	m.Ctx.SetCookie("auth", "")
	m.Redirect(beego.AppConfig.String("adminurl")+"/login", 302)
}

// Profile 管理员信息修改
func (m *IndexController) Profile() {

}
