package admin

import (
	"fmt"
)

type MainHandler struct {
	BaseHandler
}

func (m *MainHandler) Get() {
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "astaxie@gmail.com"
	m.TplNames = "index.tpl"
	fmt.Println("test main")
}
