package admin

import (
	"fmt"
	"numa/models"
	"strings"
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
		if admin.Read("account") != nil || admin.Password != models.Md5([]byte(password)) {
			m.Data["errmsg"] = "账号或密码填写错误"
		} else {
			admin.LastIp = m.GetClientIp()
            admin.LastTime = m.GetTime()
            admin.Update()
            
            
            if remember == "yes" {
                
            }
		}

	}
	
	m.TplNames = "admin/login.html"

}
