package admin

import (
	"encoding/base64"
	"fmt"
	"numa/models"
	"numa/toolkit"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ProfileController 后台设置
type ProfileController struct {
	BaseController
}

// Login 登录处理
func (m *ProfileController) Login() {

	if m.Ctx.Request.Method == "GET" {

	} else if m.Ctx.Request.Method == "POST" {

		account := strings.TrimSpace(m.GetString("account"))
		password := strings.TrimSpace(m.GetString("password"))
		remember := m.GetString("remember")

		// fmt.Println("account = " + account)
		// fmt.Println("password = " + password)

		if account != "" && password != "" {
			var admin models.Admin
			admin.Account = account

			// if account == "admin" && password == "123456" {
			// 	admin.NickName = "我是管理员"
			// 	admin.LastIP = m.GetClientIP()
			// 	admin.Password = toolkit.SHA256([]byte(password))
			// 	admin.LastTime = toolkit.GetTime()
			// 	admin.Token = toolkit.GenUID()
			// 	admin.Email = "lurma@qq.com"
			// 	admin.Status = 1
			// 	admin.Head = "default.png"
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

				admin.LastIP = m.GetClientIP()
				admin.LastTime = toolkit.GetTime()
				admin.Token = token
				admin.Update()

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
	}

	m.TplName = beego.AppConfig.String("adminpath") + "/login.html"
}

// Logout 注销登录
func (m *ProfileController) Logout() {
	m.Ctx.SetCookie("auth", "")
	m.Redirect(beego.AppConfig.String("adminurl")+"/login", 302)
}

// Profile 查看并修改管理员信息
func (m *ProfileController) Profile() {

	admin := models.Admin{Account: m.userName}
	o := orm.NewOrm()
	if err := o.Read(&admin, "Account"); err != nil {
		m.error(err.Error())
	}

	fmt.Printf("head = %s, method = %s\n", admin.Head, m.Ctx.Request.Method)

	if m.Ctx.Request.Method == "POST" {

		nickname := strings.TrimSpace(m.GetString("nickname"))
		email := strings.TrimSpace(m.GetString("email"))

		fmt.Printf("nickname = %s, email = %s\n", nickname, email)

		if m.verifyProfile(nickname, email) {

			admin.NickName = nickname
			admin.Email = email

			if num, err := o.Update(&admin, "nick_name", "email"); err == nil {
				if num > 0 {
					m.Data["profile_result"] = true
				} else {
					m.Data["profile_errmsg"] = "更新失败"
				}
			} else {
				m.error(err.Error())
				m.Data["profile_errmsg"] = "更新失败"
			}
		}

	}

	format := admin.LastTime.Format("2006-01-02 15:04:05")
	m.Data["time"] = format
	m.Data["head"] = beego.AppConfig.String("headpath") + "/" + admin.Head
	m.Data["admin"] = admin

	m.display("profile")
}

// verifyProfile 验证昵称和邮箱是否输入正确
func (m *ProfileController) verifyProfile(nickName, email string) bool {
	if nickName == "" {
		m.Data["profile_errmsg"] = "请输入昵称"
		return false
	} else if len(nickName) > 15 {
		m.Data["profile_errmsg"] = "昵称长度不能超过15个字符"
		return false
	}

	if email == "" {
		m.Data["profile_errmsg"] = "请输入邮箱地址"
		return false
	}

	if !toolkit.VerifyEmail(email) {
		m.Data["profile_errmsg"] = "邮箱填写错误"
		return false
	}

	return true
}

// Size 获取文件大小的接口
type Size interface {
	Size() int64
}

// Head 上传头像
func (m *ProfileController) Head() {

	admin := models.Admin{Account: m.userName}
	o := orm.NewOrm()
	if err := o.Read(&admin, "Account"); err != nil {
		m.error(err.Error())
	}

	if m.Ctx.Request.Method == "POST" {
		fmt.Printf("upload head...\n")
		file, handler, err := m.GetFile("head")

		defer file.Close()

		if err != nil {
			m.Data["head_errmsg"] = "上传文件错误: " + err.Error()
		} else {
			if fileSize, ok := file.(Size); ok {
				size := float64(fileSize.Size()) / (1024 * 1024)
				fmt.Printf("filesize = %fMB\n", toolkit.Round(size, 3))

				// if size > 5 {
				// 	m.Data["head_errmsg"] = "上传文件错误: 文件大小超出5M"
				// } else {

				// 能够获取到正确的文件，保存该文件
				fmt.Printf("filename = %s\n", handler.Filename)

				_, fileSuffix := toolkit.SplitFileNameAndSuffix(handler.Filename)
				fileName := toolkit.GenUID() + fileSuffix
				savePath := filepath.Join(beego.AppConfig.String("headpath"), fileName)
				fmt.Printf("filepath = %s\n", savePath)
				// 保存文件时需要相对路径，即 ./static/upload/head/
				err = m.SaveToFile("head", filepath.Join(".", savePath))
				if err != nil {
					m.Data["head_errmsg"] = "上传文件错误: " + err.Error()
				} else {
					admin.Head = fileName
					if num, err := o.Update(&admin, "head"); err == nil {
						if num > 0 {
							m.Data["head_result"] = true
						} else {
							m.Data["head_errmsg"] = "上传文件错误: 服务器保存文件失败"
						}
					} else {
						m.Data["profile_errmsg"] = "上传文件错误: 服务器保存文件失败"
					}
				}
				// }
			} else {
				m.Data["head_errmsg"] = "获取上传文件错误: 无法获取文件大小"
			}
		}
	}

	format := admin.LastTime.Format("2006-01-02 15:04:05")
	m.Data["time"] = format
	m.Data["head"] = beego.AppConfig.String("headpath") + "/" + admin.Head
	m.Data["admin"] = admin
	m.display("profile")
}

// Password 修改登录密码
func (m *ProfileController) Password() {

	admin := models.Admin{Account: m.userName}
	if err := orm.NewOrm().Read(&admin, "Account"); err != nil {
		m.error(err.Error())
	}

	if m.Ctx.Request.Method == "POST" {

		var errmsg string
		oldPassword := strings.TrimSpace(m.GetString("old-password"))
		newPassword := strings.TrimSpace(m.GetString("new-password"))
		confirmPassword := strings.TrimSpace(m.GetString("confirm-password"))

		if newPassword != "" {
			if oldPassword == "" || toolkit.SHA256([]byte(oldPassword)) != admin.Password {
				errmsg = "原密码输入错误"
			} else if len(newPassword) < 6 {
				errmsg = "密码长度不能少于6个字符"
			} else if newPassword != confirmPassword {
				errmsg = "两次输入的密码不一致"
			}

			if len(errmsg) == 0 {
				admin.Password = toolkit.SHA256([]byte(newPassword))
				admin.Update("password")
				m.Data["result"] = true
			} else {
				m.Data["errmsg"] = errmsg
			}
		}

	}

	m.Data["admin"] = admin
	m.display("password")
}
