package admin

import (
	"fmt"
	"numa/models"
	"strings"
	"github.com/astaxie/beego"
)

// LoginController 登录
type LoginController struct {
	BaseController
}

// Index Get请求
func (m *LoginController) Index() {
	m.TplNames = "admin/login.html"
}


// Login Post请求 登录处理
func (m *LoginController) Login() {

	account := strings.TrimSpace(m.GetString("account"))
	password := strings.TrimSpace(m.GetString("password"))
    
	fmt.Println("account = " + account)
    fmt.Println("password = " + password)
	
	remember := m.GetString("remember")

	if account != "" && password != "" {
		var admin models.Admin
		admin.Account = account
		
		if account == "admin" && password == "123456" {
			
			// admin.LastIp = m.GetClientIp()
            // admin.LastTime = m.GetTime()
            // admin.Update()
            
			uuid := m.GenUid()
			
			fmt.Println("ip = " + m.GetClientIp())
			fmt.Println("uuid = " + uuid)
			
			authKey := models.Md5([]byte(m.GetClientIp() + "|" + account + "|" + uuid ))
			if remember == "yes" {
                m.Ctx.SetCookie("auth", authKey, 7 * 86400)
            } else {
				m.Ctx.SetCookie("auth", authKey)
			}
			
            // m.SetSession("account", account)
			m.Redirect(beego.AppConfig.String("adminurl"), 302)
			
			return
		}
		
		if admin.Read("account") != nil || admin.Password != models.Md5([]byte(password)) {
			fmt.Println("账号或密码填写错误")
			m.Data["errmsg"] = "账号或密码填写错误"
		} else {
			admin.LastIP = m.GetClientIp()
            admin.LastTime = m.GetTime()
            admin.Update()
            
			if remember == "yes" {
                
            }
			
            m.SetSession("account", account)
			m.Redirect(beego.AppConfig.String("adminurl"), 302)   
		}
	}
	
	m.TplNames = "admin/login.html"

}
