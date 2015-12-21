package home

import (
	"github.com/astaxie/beego"
)

var(
	sess_username string
	sess_uid int64
	sess_role int64
	sess_email string
	sess_content string
	Counter map[string]string
)


type BaseHandler struct {
	beego.Controller
}


func (this *BaseHandler) Prepare() {
	// 从session里读取登录信息
	sess_username, _ = this.GetSession("username").(string)
	sess_uid, _ = this.GetSession("userid").(int64)
	sess_role, _ = this.GetSession("userrole").(int64)
	sess_email, _ = this.GetSession("useremail").(string)
	sess_content, _ = this.GetSession("usercontent").(string)

	this.Data["userid"] = sess_uid
	this.Data["username"] = sess_username
	this.Data["userrole"] = sess_role
	this.Data["useremail"] = sess_email
	this.Data["usercontent"] = sess_content

}

