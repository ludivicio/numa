package admin

import (
	"encoding/base64"
	"fmt"
	"numa/models"
	"numa/toolkit"
	"strings"
	"net"
	"github.com/astaxie/beego"
)

const (
	BigImagePath   = "./static/upload/bimage/"
	SmallImagePath = "./static/upload/simage/"
	FilePath       = "./static/upload/attachment/"
)

var pathArr []string = []string{"", BigImagePath, SmallImagePath, FilePath}

type BaseController struct {
	beego.Controller
	token          string
	userName       string
	moduleName     string
	controllerName string
	actionName     string
}

// Prepare function
func (m *BaseController) Prepare() {
	controllerName, actionName := m.GetControllerAndAction()
	fmt.Println("controllerName = " + controllerName + ", actionName = " + actionName)
	m.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	m.actionName = strings.ToLower(actionName)
	fmt.Println("controllerName = " + m.controllerName + ", actionName = " + m.actionName)
	m.auth()
}

func (m *BaseController) auth() {
	if m.controllerName == "index" && (m.actionName == "gologin" || m.actionName == "login" || m.actionName == "logout") {
		fmt.Println("login or logout...")
	} else {
		b64Auth := m.Ctx.GetCookie("auth")
		// fmt.Println("auth = " + b64Auth)
		if b64Auth != "" {
			data, err := base64.StdEncoding.DecodeString(b64Auth)
			if err == nil {
				decodeData, err := toolkit.AesDecrypt(data, []byte(beego.AppConfig.String("aeskey")))
				if err == nil {
					decodeAuth := string(decodeData)
					// fmt.Println("decode auth = " + decodeAuth)
					arr := strings.Split(decodeAuth, "|")
					if len(arr) == 2 {
						ip, token := arr[0], arr[1]
						if ip == m.GetClientIP() {
							// 先从Session中读取admin信息，如果没有的话再从数据库中读取
							// 这里先直接从数据库中读取
							var admin models.Admin
							admin.Token = token
							if admin.Read("token") == nil {
								m.token = token
								m.userName = admin.Account
							}
						} 
					}
				}
			}
		}

		if m.token == "" {
			m.Ctx.SetCookie("auth", "")
			m.Redirect(beego.AppConfig.String("adminurl")+"/login", 302)
		}
	}
}

// display 渲染模版
func (m *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = beego.AppConfig.String("adminpath") + "/" + tpl[0] + ".html"
	} else {
		tplname = beego.AppConfig.String("adminpath") + "/" + m.controllerName + ".html"
	}
	
	fmt.Println("contrName = " + m.controllerName)
	
	m.Data["contrName"] = m.controllerName
	m.Data["adminUrl"] = beego.AppConfig.String("adminurl")
	m.Data["userName"] = m.userName
	m.Layout = beego.AppConfig.String("adminpath") + "/layout.html"
	m.TplNames = tplname
}

//error 显示错误
func (m *BaseController) error(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, m.Ctx.Request.Referer())
	}
	m.Data["contrName"] = m.controllerName
	m.Data["adminUrl"] = beego.AppConfig.String("adminurl")
	m.Data["userName"] = m.userName
	m.Data["msg"] = msg[0]
	m.Data["redirect"] = msg[1]
	m.Layout = beego.AppConfig.String("adminpath") + "/layout.html"
	m.TplNames = beego.AppConfig.String("adminpath") + "/" + "error.html"
	m.Render()
	m.StopRun()
}

// GetClientIP 获取用户IP地址
func (m *BaseController) GetClientIP() string {
	// return m.Ctx.Request.Header.Get("x-forwarded-for")
	ip, _, _ := net.SplitHostPort(m.Ctx.Request.RemoteAddr)
	return ip
}

// IsPost 是否post提交
func (m *BaseController) IsPost() bool {
	return m.Ctx.Request.Method == "POST"
}
